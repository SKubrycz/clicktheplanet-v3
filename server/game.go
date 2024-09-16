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
	Name          string `json:"name"`
	CurrentHealth string `json:"currentHealth"`
	HealthPercent int    `json:"healthPercent"`
	MaxHealth     string `json:"maxHealth"`
	CurrentLevel  int    `json:"currentLevel"`
	CurrentStage  int    `json:"currentStage"`
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
			Name:          g.Planet.Name,
			CurrentHealth: g.Planet.CurrentHealth.String(),
			HealthPercent: percent,
			MaxHealth:     g.Planet.MaxHealth.String(),
			CurrentLevel:  g.CurrentLevel,
			CurrentStage:  g.CurrentStage,
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
