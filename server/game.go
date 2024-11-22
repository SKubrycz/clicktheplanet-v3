package main

import (
	"encoding/json"
	"fmt"
)

// Game logic here
// ActionHandler could be written here for each allowed user action
// Actions:
// - click
// - buy store upgrade {id}
// - buy ship upgrade {id}
// - previous level
// - next level

type UserClick struct {
	Gold                    string                            `json:"gold"`
	Diamonds                int64                             `json:"diamonds"`
	CurrentDamage           string                            `json:"currentDamage"`
	MaxDamage               string                            `json:"maxDamage"`
	DamageDone              DamageDoneData                    `json:"damageDone"`
	PlanetName              string                            `json:"planetName"`
	CurrentHealth           string                            `json:"currentHealth"`
	HealthPercent           int                               `json:"healthPercent"`
	MaxHealth               string                            `json:"maxHealth"`
	PlanetGold              string                            `json:"planetGold"`
	IsBoss                  bool                              `json:"isBoss"`
	DiamondPlanet           DiamondPlanetDataMessage          `json:"diamondPlanet"`
	CurrentLevel            int64                             `json:"currentLevel"`
	MaxLevel                int64                             `json:"maxLevel"`
	CurrentStage            uint8                             `json:"currentStage"`
	MaxStage                uint8                             `json:"maxStage"`
	DiamondUpgradesUnlocked bool                              `json:"diamondUpgradesUnlocked"`
	PlanetsDestroyed        string                            `json:"planetsDestroyed"`
	Store                   map[int]StoreDataMessage          `json:"store"`
	Ship                    map[int]ShipDataMessage           `json:"ship"`
	DiamondUpgrade          map[int]DiamondUpgradeDataMessage `json:"diamondUpgrade"`
}

type ActionMessage struct {
	Action string `json:"action"`
	Data   any    `json:"data"`
}

type DamageDoneData struct {
	Damage   string `json:"damage"`
	Critical bool   `json:"critical"`
}

type UpgradeMessage struct {
	Upgrade string `json:"upgrade"`
	Index   int    `json:"index"`
	Levels  int    `json:"levels"`
}

type StoreDataMessage struct {
	Index  int    `json:"index"`
	Level  int64  `json:"level"`
	Cost   string `json:"cost"`
	Damage string `json:"damage"`
	Locked bool   `json:"locked"`
}

type StoreDataMessageWrapper struct {
	Gold     string                   `json:"gold"`
	Diamonds int64                    `json:"diamonds"`
	Store    map[int]StoreDataMessage `json:"store"`
}

type DiamondUpgradeDataMessage struct {
	Index      int    `json:"index"`
	Level      int64  `json:"level"`
	Multiplier string `json:"multiplier"`
	Cost       int64  `json:"cost"`
}

type ShipDataMessage struct {
	Index      int     `json:"index"`
	Level      int64   `json:"level"`
	Cost       string  `json:"cost"`
	Multiplier float64 `json:"multiplier"`
	Damage     string  `json:"damage"`
	Locked     bool    `json:"locked"`
}

type ShipDataMessageWrapper struct {
	Gold     string          `json:"gold"`
	Diamonds int64           `json:"diamonds"`
	Ship     ShipDataMessage `json:"ship"`
}

type UpgradeDataMessageWrapper struct {
	Gold           string                            `json:"gold"`
	Diamonds       int64                             `json:"diamonds"`
	CurrentDamage  string                            `json:"currentDamage"`
	MaxDamage      string                            `json:"maxDamage"`
	PlanetGold     string                            `json:"planetGold"`
	Store          map[int]StoreDataMessage          `json:"store"`
	Ship           map[int]ShipDataMessage           `json:"ship"`
	DiamondUpgrade map[int]DiamondUpgradeDataMessage `json:"diamondUpgrade"`
}

type LevelDataMessage struct {
	CurrentLevel  int64  `json:"currentLevel"`
	CurrentStage  uint8  `json:"currentStage"`
	PlanetName    string `json:"planetName"`
	CurrentHealth string `json:"currentHealth"`
	HealthPercent int    `json:"healthPercent"`
	MaxHealth     string `json:"maxHealth"`
	PlanetGold    string `json:"planetGold"`
	IsBoss        bool   `json:"isBoss"`
}

type DiamondPlanetDataMessage struct {
	Diamonds        int     `json:"diamonds"`
	Chance          float64 `json:"chance"`
	IsDiamondPlanet bool    `json:"isDiamondPlanet"`
}

type ErrorDataMessage struct {
	Error string `json:"error"`
}

func ActionHandler(g *Game, action string) []byte {
	fmt.Println("received message: ", []byte(action))
	// Note: JSON.Stringify() from the frontend
	// adds additional quotes around the string,
	// so "click" instead of plain click
	if action == "click" {
		g.ClickThePlanet(g.CurrentDamage, true)
		damageDone := DamageDoneData{
			Damage:   g.DisplayNumber(g.DamageDone.Damage),
			Critical: g.DamageDone.Critical,
		}
		percent := g.GetHealthPercent()

		diamondPlanet := DiamondPlanetDataMessage{
			Diamonds:        int(g.Planet.DiamondPlanet.Diamonds),
			Chance:          g.Planet.DiamondPlanet.Chance,
			IsDiamondPlanet: g.Planet.DiamondPlanet.IsDiamondPlanet,
		}

		store := map[int]StoreDataMessage{}
		for k := range g.Store {
			s := new(StoreDataMessage)
			s.Level = g.Store[k].Level
			s.Cost = g.DisplayNumber(g.Store[k].Cost)
			s.Damage = g.DisplayNumber(g.Store[k].Damage)
			s.Locked = g.Store[k].Locked
			store[k] = *s
		}
		ship := map[int]ShipDataMessage{}
		for k := range g.Ship {
			s := new(ShipDataMessage)
			s.Level = g.Ship[k].Level
			s.Cost = g.DisplayNumber(g.Ship[k].Cost)
			s.Multiplier = g.Ship[k].Multiplier
			s.Locked = g.Ship[k].Locked
			if k == 1 {
				s.Damage = g.DisplayNumber(g.Dps)
			} else {
				s.Damage = g.DisplayNumber(g.Ship[k].Damage)
			}
			ship[k] = *s
		}
		diamondUpgrade := map[int]DiamondUpgradeDataMessage{}
		for k := range g.DiamondUpgrade {
			d := new(DiamondUpgradeDataMessage)
			d.Index = k
			d.Level = g.DiamondUpgrade[k].Level
			d.Cost = g.DiamondUpgrade[k].Cost
			d.Multiplier = g.DisplayNumber(g.DiamondUpgrade[k].Multiplier)
			diamondUpgrade[k] = *d
		}
		// Wrapper to indicate for frontend how to behave depending on the action field
		message := ActionMessage{
			Action: action,
			Data: UserClick{
				Gold:                    g.DisplayNumber(g.Gold),
				Diamonds:                g.Diamonds,
				CurrentDamage:           g.DisplayNumber(g.CurrentDamage),
				MaxDamage:               g.DisplayNumber(g.MaxDamage),
				DamageDone:              damageDone,
				PlanetName:              g.Planet.Name,
				CurrentHealth:           g.DisplayNumber(g.Planet.CurrentHealth),
				HealthPercent:           percent,
				MaxHealth:               g.DisplayNumber(g.Planet.MaxHealth),
				PlanetGold:              g.DisplayNumber(g.Planet.Gold),
				IsBoss:                  g.Planet.IsBoss,
				DiamondPlanet:           diamondPlanet,
				CurrentLevel:            g.CurrentLevel,
				MaxLevel:                g.MaxLevel,
				CurrentStage:            g.CurrentStage,
				MaxStage:                g.MaxStage,
				DiamondUpgradesUnlocked: g.DiamondUpgradesUnlocked,
				PlanetsDestroyed:        g.DisplayNumber(g.PlanetsDestroyed),
				Store:                   store,
				Ship:                    ship,
				DiamondUpgrade:          diamondUpgrade,
			},
		}
		encoded, _ := json.Marshal(message)
		return []byte(encoded)
	} else if action == "init" {
		g.CheckBoss()
		g.CalculatePlanetHealth()
		g.GeneratePlanetName()
		g.CheckDiamondUpgradeUnlock()
		g.CalculateDiamondUpgrade(-1)
		g.CalculateCurrentDamage()
		g.CalculateStore(-1)
		g.CalculateShip(-1)

		g.DamageDone.Damage = g.CurrentDamage
		g.DamageDone.Critical = false

		damageDone := DamageDoneData{
			Damage:   g.DisplayNumber(g.DamageDone.Damage),
			Critical: g.DamageDone.Critical,
		}

		diamondPlanet := DiamondPlanetDataMessage{
			Diamonds:        int(g.Planet.DiamondPlanet.Diamonds),
			Chance:          g.Planet.DiamondPlanet.Chance,
			IsDiamondPlanet: g.Planet.DiamondPlanet.IsDiamondPlanet,
		}

		store := map[int]StoreDataMessage{}
		for k := range g.Store {
			s := new(StoreDataMessage)
			s.Index = k
			s.Level = g.Store[k].Level
			s.Cost = g.DisplayNumber(g.Store[k].Cost)
			s.Damage = g.DisplayNumber(g.Store[k].Damage)
			s.Locked = g.Store[k].Locked
			store[k] = *s
		}
		ship := map[int]ShipDataMessage{}
		for k := range g.Ship {
			s := new(ShipDataMessage)
			s.Index = k
			s.Level = g.Ship[k].Level
			s.Cost = g.DisplayNumber(g.Ship[k].Cost)
			s.Multiplier = g.Ship[k].Multiplier
			s.Locked = g.Ship[k].Locked
			if k == 1 {
				s.Damage = g.DisplayNumber(g.Dps)
			} else {
				s.Damage = g.DisplayNumber(g.Ship[k].Damage)
			}
			ship[k] = *s
		}
		diamondUpgrade := map[int]DiamondUpgradeDataMessage{}
		for k := range g.DiamondUpgrade {
			d := new(DiamondUpgradeDataMessage)
			d.Index = k
			d.Level = g.DiamondUpgrade[k].Level
			d.Cost = g.DiamondUpgrade[k].Cost
			d.Multiplier = g.DisplayNumber(g.DiamondUpgrade[k].Multiplier)
			diamondUpgrade[k] = *d
		}
		percent := g.GetHealthPercent()
		g.ConvertNumber(g.Planet.MaxHealth, g.Planet.MaxHealth)
		message := ActionMessage{
			Action: action,
			Data: UserClick{
				Gold:                    g.DisplayNumber(g.Gold),
				Diamonds:                g.Diamonds,
				CurrentDamage:           g.DisplayNumber(g.CurrentDamage),
				MaxDamage:               g.DisplayNumber(g.MaxDamage),
				DamageDone:              damageDone,
				PlanetName:              g.Planet.Name,
				CurrentHealth:           g.DisplayNumber(g.Planet.CurrentHealth),
				HealthPercent:           percent,
				MaxHealth:               g.DisplayNumber(g.Planet.MaxHealth),
				PlanetGold:              g.DisplayNumber(g.Planet.Gold),
				IsBoss:                  g.Planet.IsBoss,
				DiamondPlanet:           diamondPlanet,
				CurrentLevel:            g.CurrentLevel,
				MaxLevel:                g.MaxLevel,
				CurrentStage:            g.CurrentStage,
				MaxStage:                g.MaxStage,
				DiamondUpgradesUnlocked: g.DiamondUpgradesUnlocked,
				PlanetsDestroyed:        g.DisplayNumber(g.PlanetsDestroyed),
				Store:                   store,
				Ship:                    ship,
				DiamondUpgrade:          diamondUpgrade,
			},
		}
		encoded, _ := json.Marshal(message)
		return []byte(encoded)
	} else if action == "previous" || action == "next" {
		// Previous level or next
		if action == "previous" {
			g.PreviousLevel()
		}
		if action == "next" {
			g.NextLevel()
		}
		percent := g.GetHealthPercent()

		message := LevelDataMessage{
			CurrentLevel:  g.CurrentLevel,
			CurrentStage:  g.CurrentStage,
			PlanetName:    g.Planet.Name,
			CurrentHealth: g.DisplayNumber(g.Planet.CurrentHealth),
			HealthPercent: percent,
			MaxHealth:     g.DisplayNumber(g.Planet.MaxHealth),
			PlanetGold:    g.DisplayNumber(g.Planet.Gold),
			IsBoss:        g.Planet.IsBoss,
		}
		enc, _ := json.Marshal(message)
		return []byte(enc)
	} else {
		unmarshaled := new(UpgradeMessage)
		if action != "dps" {
			err := json.Unmarshal([]byte(action), unmarshaled)
			if err != nil {
				fmt.Println("ERROR json ActionHandler: ", err)
			}
		}
		if unmarshaled.Upgrade == "store" || unmarshaled.Upgrade == "ship" || unmarshaled.Upgrade == "diamond-upgrade" {
			if unmarshaled.Upgrade == "store" {
				err := g.UpgradeStore(unmarshaled.Index, unmarshaled.Levels)
				if err != "" {
					message := ActionMessage{
						Action: "error",
						Data: ErrorDataMessage{
							Error: err,
						},
					}
					enc, _ := json.Marshal(message)
					return []byte(enc)
				}
			}
			if unmarshaled.Upgrade == "ship" {
				err := g.UpgradeShip(unmarshaled.Index, unmarshaled.Levels)
				if err != "" {
					message := ActionMessage{
						Action: "error",
						Data: ErrorDataMessage{
							Error: err,
						},
					}
					enc, _ := json.Marshal(message)
					return []byte(enc)
				}
			}
			if unmarshaled.Upgrade == "diamond-upgrade" {
				err := g.UpgradeDiamondUpgrade(unmarshaled.Index, unmarshaled.Levels)
				if err != "" {
					message := ActionMessage{
						Action: "error",
						Data: ErrorDataMessage{
							Error: err,
						},
					}
					enc, _ := json.Marshal(message)
					return []byte(enc)
				}
			}
			g.CalculateDiamondUpgrade(unmarshaled.Index)
			g.CalculateStore(unmarshaled.Index)
			g.CalculateShip(-1)

			store := map[int]StoreDataMessage{}
			for k := range g.Store {
				s := new(StoreDataMessage)
				s.Index = k
				s.Level = g.Store[k].Level
				s.Cost = g.DisplayNumber(g.Store[k].Cost)
				s.Damage = g.DisplayNumber(g.Store[k].Damage)
				s.Locked = g.Store[k].Locked
				store[k] = *s
			}
			ship := map[int]ShipDataMessage{}
			for k := range g.Ship {
				s := new(ShipDataMessage)
				s.Index = k
				s.Level = g.Ship[k].Level
				s.Cost = g.DisplayNumber(g.Ship[k].Cost)
				s.Multiplier = g.Ship[k].Multiplier
				s.Locked = g.Ship[k].Locked
				if k == 1 {
					s.Damage = g.DisplayNumber(g.Dps)
				} else {
					s.Damage = g.DisplayNumber(g.Ship[k].Damage)
				}
				ship[k] = *s
			}
			diamondUpgrade := map[int]DiamondUpgradeDataMessage{}
			for k := range g.DiamondUpgrade {
				d := new(DiamondUpgradeDataMessage)
				d.Index = k
				d.Level = g.DiamondUpgrade[k].Level
				d.Cost = g.DiamondUpgrade[k].Cost
				d.Multiplier = g.DisplayNumber(g.DiamondUpgrade[k].Multiplier)
				diamondUpgrade[k] = *d
			}
			message := ActionMessage{
				Action: "upgrade",
				Data: UpgradeDataMessageWrapper{
					Gold:           g.DisplayNumber(g.Gold),
					Diamonds:       g.Diamonds,
					CurrentDamage:  g.DisplayNumber(g.CurrentDamage),
					MaxDamage:      g.DisplayNumber(g.MaxDamage),
					PlanetGold:     g.DisplayNumber(g.Planet.Gold),
					Store:          store,
					Ship:           ship,
					DiamondUpgrade: diamondUpgrade,
				},
			}
			encoded, _ := json.Marshal(message)
			return []byte(encoded)
		}
	}
	// ^ switch case to be considered

	message := ActionMessage{
		Action: "unknown",
		Data:   nil,
	}
	encoded, _ := json.Marshal(message)
	return []byte(encoded)
}

func DealDps(g *Game) []byte {
	g.ClickThePlanet(g.Ship[1].Damage, false)
	percent := g.GetHealthPercent()

	diamondPlanet := DiamondPlanetDataMessage{
		Diamonds:        int(g.Planet.DiamondPlanet.Diamonds),
		Chance:          g.Planet.DiamondPlanet.Chance,
		IsDiamondPlanet: g.Planet.DiamondPlanet.IsDiamondPlanet,
	}

	store := map[int]StoreDataMessage{}
	for k := range g.Store {
		s := new(StoreDataMessage)
		s.Index = k
		s.Level = g.Store[k].Level
		s.Cost = g.DisplayNumber(g.Store[k].Cost)
		s.Damage = g.DisplayNumber(g.Store[k].Damage)
		s.Locked = g.Store[k].Locked
		store[k] = *s
	}
	ship := map[int]ShipDataMessage{}
	for k := range g.Ship {
		s := new(ShipDataMessage)
		s.Index = k
		s.Level = g.Ship[k].Level
		s.Cost = g.DisplayNumber(g.Ship[k].Cost)
		s.Multiplier = g.Ship[k].Multiplier
		s.Locked = g.Ship[k].Locked
		if k == 1 {
			s.Damage = g.DisplayNumber(g.Dps)
		} else {
			s.Damage = g.DisplayNumber(g.Ship[k].Damage)
		}
		ship[k] = *s
	}
	diamondUpgrade := map[int]DiamondUpgradeDataMessage{}
	for k := range g.DiamondUpgrade {
		d := new(DiamondUpgradeDataMessage)
		d.Index = k
		d.Level = g.DiamondUpgrade[k].Level
		d.Cost = g.DiamondUpgrade[k].Cost
		d.Multiplier = g.DisplayNumber(g.DiamondUpgrade[k].Multiplier)
		diamondUpgrade[k] = *d
	}

	message := ActionMessage{
		Action: "dps",
		Data: UserClick{
			Gold:                    g.DisplayNumber(g.Gold),
			Diamonds:                g.Diamonds,
			CurrentDamage:           g.DisplayNumber(g.CurrentDamage),
			MaxDamage:               g.DisplayNumber(g.MaxDamage),
			PlanetName:              g.Planet.Name,
			CurrentHealth:           g.DisplayNumber(g.Planet.CurrentHealth),
			HealthPercent:           percent,
			MaxHealth:               g.DisplayNumber(g.Planet.MaxHealth),
			PlanetGold:              g.DisplayNumber(g.Planet.Gold),
			IsBoss:                  g.Planet.IsBoss,
			DiamondPlanet:           diamondPlanet,
			CurrentLevel:            g.CurrentLevel,
			MaxLevel:                g.MaxLevel,
			CurrentStage:            g.CurrentStage,
			MaxStage:                g.MaxStage,
			DiamondUpgradesUnlocked: g.DiamondUpgradesUnlocked,
			PlanetsDestroyed:        g.DisplayNumber(g.PlanetsDestroyed),
			Store:                   store,
			Ship:                    ship,
			DiamondUpgrade:          diamondUpgrade,
		},
	}
	enc, _ := json.Marshal(message)
	return []byte(enc)
}
