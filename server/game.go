package main

import (
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

func ActionHandler(g *Game, p string) []byte {
	fmt.Println("received message: ", []byte(p))
	// Note: JSON.Stringify() from the frontend
	// adds additional quotes around the string,
	// so "click" instead of plain click
	if p == "click" {
		g.ClickThePlanet()
		return []byte(g.Planet.CurrentHealth.String())
	} else {
		return []byte("nothing")
	}
	// ^ switch case to be considered
}
