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
	Gold             string `json:"gold"`
	Diamonds         int64  `json:"diamonds"`
	CurrentDamage    string `json:"currentDamage"`
	MaxDamage        string `json:"maxDamage"`
	PlanetName       string `json:"planetName"`
	CurrentHealth    string `json:"currentHealth"`
	HealthPercent    int    `json:"healthPercent"`
	MaxHealth        string `json:"maxHealth"`
	CurrentLevel     int64  `json:"currentLevel"`
	MaxLevel         int64  `json:"maxLevel"`
	CurrentStage     uint8  `json:"currentStage"`
	MaxStage         uint8  `json:"maxStage"`
	PlanetsDestroyed string `json:"planetsDestroyed"`
}

type ActionMessage struct {
	Message string
}

type UpgradeMessage struct {
	Upgrade string `json:"upgrade"`
	Index   int    `json:"index"`
}

func ActionHandler(g *Game, action string) []byte {
	fmt.Println("received message: ", []byte(action))
	// Note: JSON.Stringify() from the frontend
	// adds additional quotes around the string,
	// so "click" instead of plain click
	if action == "click" {
		g.ClickThePlanet()
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
		}
		encoded, _ := json.Marshal(userClick)
		return []byte(encoded)
	} else if action == "init" {
		g.CalculatePlanetHealth()
		//g.CalculateStore
		//g.CalculateShip
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
		}
		encoded, _ := json.Marshal(userClick)
		return []byte(encoded)
	} else {
		unmarshaled := new(UpgradeMessage)
		err := json.Unmarshal([]byte(action), unmarshaled)
		if err != nil {
			fmt.Println("ERROR json ActionHandler: ", err)
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
