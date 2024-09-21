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
	Gold             string                   `json:"gold"`
	Diamonds         int64                    `json:"diamonds"`
	CurrentDamage    string                   `json:"currentDamage"`
	MaxDamage        string                   `json:"maxDamage"`
	PlanetName       string                   `json:"planetName"`
	CurrentHealth    string                   `json:"currentHealth"`
	HealthPercent    int                      `json:"healthPercent"`
	MaxHealth        string                   `json:"maxHealth"`
	CurrentLevel     int64                    `json:"currentLevel"`
	MaxLevel         int64                    `json:"maxLevel"`
	CurrentStage     uint8                    `json:"currentStage"`
	MaxStage         uint8                    `json:"maxStage"`
	PlanetsDestroyed string                   `json:"planetsDestroyed"`
	Store            map[int]StoreDataMessage `json:"store"`
	Ship             map[int]ShipDataMessage  `json:"ship"`
}

type ActionMessage struct {
	Action string `json:"action"`
	Data   any    `json:"data"`
}

type UpgradeMessage struct {
	Upgrade string `json:"upgrade"`
	Index   int    `json:"index"`
}

type StoreDataMessage struct {
	Level  int64  `json:"level"`
	Cost   string `json:"cost"`
	Damage string `json:"damage"`
}

type StoreDataMessageWrapper struct {
	Store StoreDataMessage `json:"store"`
}

type ShipDataMessage struct {
	Level      int64   `json:"level"`
	Cost       string  `json:"cost"`
	Multiplier float64 `json:"multiplier"`
	Damage     string  `json:"damage"`
}

type ShipDataMessageWrapper struct {
	Ship ShipDataMessage `json:"ship"`
}

func ActionHandler(g *Game, action string) []byte {
	fmt.Println("received message: ", []byte(action))
	// Note: JSON.Stringify() from the frontend
	// adds additional quotes around the string,
	// so "click" instead of plain click
	if action == "click" {
		g.ClickThePlanet()
		percent := g.GetHealthPercent()
		store := map[int]StoreDataMessage{}
		for k := range g.Store {
			s := new(StoreDataMessage)
			s.Level = g.Store[k].Level
			s.Cost = g.Store[k].Cost.String()
			s.Damage = g.Store[k].Damage.String()
			store[k] = *s
		}
		ship := map[int]ShipDataMessage{}
		for k := range g.Store {
			s := new(ShipDataMessage)
			s.Level = g.Ship[k].Level
			s.Cost = g.Ship[k].Cost.String()
			s.Damage = g.Ship[k].Damage.String()
			ship[k] = *s
		}
		// Wrapper to indicate for frontend how to behave depending on the action field
		message := ActionMessage{
			Action: action,
			Data: UserClick{
				Gold:             g.Gold.String(),
				Diamonds:         g.Diamonds,
				CurrentDamage:    g.CurrentDamage.String(),
				MaxDamage:        g.MaxDamage.String(),
				PlanetName:       g.Planet.Name,
				CurrentHealth:    g.Planet.CurrentHealth.String(),
				HealthPercent:    percent,
				MaxHealth:        g.Planet.MaxHealth.String(),
				CurrentLevel:     g.CurrentLevel,
				MaxLevel:         g.MaxLevel,
				CurrentStage:     g.CurrentStage,
				MaxStage:         g.MaxStage,
				PlanetsDestroyed: g.PlanetsDestroyed.String(),
				Store:            store,
				Ship:             ship,
			},
		}
		encoded, _ := json.Marshal(message)
		return []byte(encoded)
	} else if action == "init" {
		g.CalculatePlanetHealth()
		g.CalculateStore(-1)
		g.CalculateCurrentDamage()
		//g.CalculateShip
		store := map[int]StoreDataMessage{}
		for k := range g.Store {
			s := new(StoreDataMessage)
			s.Level = g.Store[k].Level
			s.Cost = g.Store[k].Cost.String()
			s.Damage = g.Store[k].Damage.String()
			store[k] = *s
		}
		ship := map[int]ShipDataMessage{}
		for k := range g.Store {
			s := new(ShipDataMessage)
			s.Level = g.Ship[k].Level
			s.Cost = g.Ship[k].Cost.String()
			s.Damage = g.Ship[k].Damage.String()
			ship[k] = *s
		}
		percent := g.GetHealthPercent()
		message := ActionMessage{
			Action: action,
			Data: UserClick{
				Gold:             g.Gold.String(),
				Diamonds:         g.Diamonds,
				CurrentDamage:    g.CurrentDamage.String(),
				MaxDamage:        g.MaxDamage.String(),
				PlanetName:       g.Planet.Name,
				CurrentHealth:    g.Planet.CurrentHealth.String(),
				HealthPercent:    percent,
				MaxHealth:        g.Planet.MaxHealth.String(),
				CurrentLevel:     g.CurrentLevel,
				MaxLevel:         g.MaxLevel,
				CurrentStage:     g.CurrentStage,
				MaxStage:         g.MaxStage,
				PlanetsDestroyed: g.PlanetsDestroyed.String(),
				Store:            store,
				Ship:             ship,
			},
		}
		encoded, _ := json.Marshal(message)
		return []byte(encoded)
	} else {
		unmarshaled := new(UpgradeMessage)
		err := json.Unmarshal([]byte(action), unmarshaled)
		if err != nil {
			fmt.Println("ERROR json ActionHandler: ", err)
		}
		if unmarshaled.Upgrade == "store" {
			g.UpgradeStore(unmarshaled.Index)
			g.CalculateStore(unmarshaled.Index)
			storeData := StoreDataMessage{
				Level:  g.Store[unmarshaled.Index].Level,
				Cost:   g.Store[unmarshaled.Index].Cost.String(),
				Damage: g.Store[unmarshaled.Index].Damage.String(),
			}
			message := ActionMessage{
				Action: "store",
				Data: StoreDataMessageWrapper{
					Store: storeData,
				},
			}
			encoded, _ := json.Marshal(message)
			return []byte(encoded)
		}
		if unmarshaled.Upgrade == "ship" {
			g.UpgradeShip(unmarshaled.Index)
			g.CalculateShip(unmarshaled.Index)
			shipData := ShipDataMessage{
				Level:      g.Ship[unmarshaled.Index].Level,
				Cost:       g.Ship[unmarshaled.Index].Cost.String(),
				Multiplier: g.Ship[unmarshaled.Index].Multiplier,
				Damage:     g.Ship[unmarshaled.Index].Damage.String(),
			}
			message := ActionMessage{
				Action: "ship",
				Data: ShipDataMessageWrapper{
					Ship: shipData,
				},
			}
			encoded, _ := json.Marshal(message)
			return []byte(encoded)
		}
		fmt.Println("Unmarshaled", unmarshaled)
	} /* else {
		message := Message{
			Message: "error sending data",
		}
		encMessage, _ := json.Marshal(message)
		return []byte(encMessage)
	} */
	// ^ switch case to be considered

	message := ActionMessage{
		Action: "unknown",
		Data:   nil,
	}
	encoded, _ := json.Marshal(message)
	return []byte(encoded)
}
