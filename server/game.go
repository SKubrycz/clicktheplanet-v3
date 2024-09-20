package main

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	Gold             string                      `json:"gold"`
	Diamonds         int64                       `json:"diamonds"`
	CurrentDamage    string                      `json:"currentDamage"`
	MaxDamage        string                      `json:"maxDamage"`
	PlanetName       string                      `json:"planetName"`
	CurrentHealth    string                      `json:"currentHealth"`
	HealthPercent    int                         `json:"healthPercent"`
	MaxHealth        string                      `json:"maxHealth"`
	CurrentLevel     int64                       `json:"currentLevel"`
	MaxLevel         int64                       `json:"maxLevel"`
	CurrentStage     uint8                       `json:"currentStage"`
	MaxStage         uint8                       `json:"maxStage"`
	PlanetsDestroyed string                      `json:"planetsDestroyed"`
	Store            map[string]StoreDataMessage `json:"store"`
	Ship             map[string]ShipDataMessage  `json:"ship"`
}

type ActionMessage struct {
	Message string
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

type ShipDataMessage struct {
	Level      int64   `json:"level"`
	Cost       string  `json:"cost"`
	Multiplier float64 `json:"multiplier"`
	Damage     string  `json:"damage"`
}

func ActionHandler(g *Game, action string) []byte {
	fmt.Println("received message: ", []byte(action))
	// Note: JSON.Stringify() from the frontend
	// adds additional quotes around the string,
	// so "click" instead of plain click
	if action == "click" {
		g.ClickThePlanet()
		percent := g.GetHealthPercent()
		store := map[string]StoreDataMessage{}
		for k := range g.Store {
			s := new(StoreDataMessage)
			s.Level = g.Store[k].Level
			s.Cost = g.Store[k].Cost.String()
			s.Damage = g.Store[k].Damage.String()
			store[k] = *s
		}
		ship := map[string]ShipDataMessage{}
		for k := range g.Store {
			s := new(ShipDataMessage)
			s.Level = g.Ship[k].Level
			s.Cost = g.Ship[k].Cost.String()
			s.Damage = g.Ship[k].Damage.String()
			ship[k] = *s
		}
		userClick := UserClick{
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
		}
		encoded, _ := json.Marshal(userClick)
		return []byte(encoded)
	} else if action == "init" {
		g.CalculatePlanetHealth()
		g.CalculateStore(-1)
		//g.CalculateShip
		store := map[string]StoreDataMessage{}
		for k := range g.Store {
			s := new(StoreDataMessage)
			s.Level = g.Store[k].Level
			s.Cost = g.Store[k].Cost.String()
			s.Damage = g.Store[k].Damage.String()
			store[k] = *s
		}
		ship := map[string]ShipDataMessage{}
		for k := range g.Store {
			s := new(ShipDataMessage)
			s.Level = g.Ship[k].Level
			s.Cost = g.Ship[k].Cost.String()
			s.Damage = g.Ship[k].Damage.String()
			ship[k] = *s
		}
		percent := g.GetHealthPercent()
		userClick := UserClick{
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
		}
		encoded, _ := json.Marshal(userClick)
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
			indexStr := strconv.Itoa(unmarshaled.Index)
			storeData := StoreDataMessage{
				Level:  g.Store[indexStr].Level,
				Cost:   g.Store[indexStr].Cost.String(),
				Damage: g.Store[indexStr].Damage.String(),
			}
			encoded, _ := json.Marshal(storeData)
			return []byte(encoded)
		}
		if unmarshaled.Upgrade == "ship" {
			g.UpgradeShip(unmarshaled.Index)
			g.CalculateShip(unmarshaled.Index)
			indexStr := strconv.Itoa(unmarshaled.Index)
			shipData := ShipDataMessage{
				Level:      g.Ship[indexStr].Level,
				Cost:       g.Ship[indexStr].Cost.String(),
				Multiplier: g.Ship[indexStr].Multiplier,
				Damage:     g.Ship[indexStr].Damage.String(),
			}
			encoded, _ := json.Marshal(shipData)
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
		Message: action,
	}
	encoded, _ := json.Marshal(message)
	return []byte(encoded)
}
