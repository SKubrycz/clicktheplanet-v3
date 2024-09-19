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
	Diamonds         int    `json:"diamonds"`
	CurrentDamage    string `json:"currentDamage"`
	MaxDamage        string `json:"maxDamage"`
	PlanetName       string `json:"planetName"`
	CurrentHealth    string `json:"currentHealth"`
	HealthPercent    int    `json:"healthPercent"`
	MaxHealth        string `json:"maxHealth"`
	CurrentLevel     int    `json:"currentLevel"`
	MaxLevel         int    `json:"maxLevel"`
	CurrentStage     int    `json:"currentStage"`
	MaxStage         int    `json:"maxStage"`
	PlanetsDestroyed string `json:"planetsDestroyed"`
}

type ActionMessage struct {
	Message string
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
