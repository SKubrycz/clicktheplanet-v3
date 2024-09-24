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
	Index  int    `json:"index"`
	Level  int64  `json:"level"`
	Cost   string `json:"cost"`
	Damage string `json:"damage"`
}

type StoreDataMessageWrapper struct {
	Gold     string           `json:"gold"`
	Diamonds int64            `json:"diamonds"`
	Store    StoreDataMessage `json:"store"`
}

type ShipDataMessage struct {
	Index      int     `json:"index"`
	Level      int64   `json:"level"`
	Cost       string  `json:"cost"`
	Multiplier float64 `json:"multiplier"`
	Damage     string  `json:"damage"`
}

type ShipDataMessageWrapper struct {
	Gold     string          `json:"gold"`
	Diamonds int64           `json:"diamonds"`
	Ship     ShipDataMessage `json:"ship"`
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
			s.Cost = g.DisplayNumber(g.Store[k].Cost)
			s.Damage = g.DisplayNumber(g.Store[k].Damage)
			store[k] = *s
		}
		ship := map[int]ShipDataMessage{}
		for k := range g.Store {
			s := new(ShipDataMessage)
			s.Level = g.Ship[k].Level
			s.Cost = g.DisplayNumber(g.Ship[k].Cost)
			s.Damage = g.DisplayNumber(g.Ship[k].Damage)
			ship[k] = *s
		}
		// Wrapper to indicate for frontend how to behave depending on the action field
		message := ActionMessage{
			Action: action,
			Data: UserClick{
				Gold:             g.DisplayNumber(g.Gold),
				Diamonds:         g.Diamonds,
				CurrentDamage:    g.DisplayNumber(g.CurrentDamage),
				MaxDamage:        g.DisplayNumber(g.MaxDamage),
				PlanetName:       g.Planet.Name,
				CurrentHealth:    g.DisplayNumber(g.Planet.CurrentHealth),
				HealthPercent:    percent,
				MaxHealth:        g.DisplayNumber(g.Planet.MaxHealth),
				CurrentLevel:     g.CurrentLevel,
				MaxLevel:         g.MaxLevel,
				CurrentStage:     g.CurrentStage,
				MaxStage:         g.MaxStage,
				PlanetsDestroyed: g.DisplayNumber(g.PlanetsDestroyed),
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
			s.Index = k
			s.Level = g.Store[k].Level
			s.Cost = g.DisplayNumber(g.Store[k].Cost)
			s.Damage = g.DisplayNumber(g.Store[k].Damage)
			store[k] = *s
		}
		ship := map[int]ShipDataMessage{}
		for k := range g.Store {
			s := new(ShipDataMessage)
			s.Index = k
			s.Level = g.Ship[k].Level
			s.Cost = g.DisplayNumber(g.Ship[k].Cost)
			s.Damage = g.DisplayNumber(g.Ship[k].Damage)
			ship[k] = *s
		}
		percent := g.GetHealthPercent()
		g.ConvertNumber(g.Planet.MaxHealth, g.Planet.MaxHealth)
		message := ActionMessage{
			Action: action,
			Data: UserClick{
				Gold:             g.DisplayNumber(g.Gold),
				Diamonds:         g.Diamonds,
				CurrentDamage:    g.DisplayNumber(g.CurrentDamage),
				MaxDamage:        g.DisplayNumber(g.MaxDamage),
				PlanetName:       g.Planet.Name,
				CurrentHealth:    g.DisplayNumber(g.Planet.CurrentHealth),
				HealthPercent:    percent,
				MaxHealth:        g.DisplayNumber(g.Planet.MaxHealth),
				CurrentLevel:     g.CurrentLevel,
				MaxLevel:         g.MaxLevel,
				CurrentStage:     g.CurrentStage,
				MaxStage:         g.MaxStage,
				PlanetsDestroyed: g.DisplayNumber(g.PlanetsDestroyed),
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
				Index:  unmarshaled.Index,
				Level:  g.Store[unmarshaled.Index].Level,
				Cost:   g.DisplayNumber(g.Store[unmarshaled.Index].Cost),
				Damage: g.DisplayNumber(g.Store[unmarshaled.Index].Damage),
			}
			message := ActionMessage{
				Action: "store",
				Data: StoreDataMessageWrapper{
					Gold:     g.DisplayNumber(g.Gold),
					Diamonds: g.Diamonds,
					Store:    storeData,
				},
			}
			encoded, _ := json.Marshal(message)
			return []byte(encoded)
		}
		if unmarshaled.Upgrade == "ship" {
			g.UpgradeShip(unmarshaled.Index)
			g.CalculateShip(unmarshaled.Index)
			shipData := ShipDataMessage{
				Index:      unmarshaled.Index,
				Level:      g.Ship[unmarshaled.Index].Level,
				Cost:       g.DisplayNumber(g.Ship[unmarshaled.Index].Cost),
				Multiplier: g.Ship[unmarshaled.Index].Multiplier,
				Damage:     g.DisplayNumber(g.Ship[unmarshaled.Index].Damage),
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
