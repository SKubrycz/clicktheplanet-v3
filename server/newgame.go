package main

import (
	"math/big"
)

func NewGame(gameData *GameData) *Game {
	// big numbers initialization and db data assignment to variables
	id := gameData.Id

	gold := new(big.Float)
	gold.SetString(gameData.Gold)

	diamonds := gameData.Diamonds

	currentDamage := new(big.Float)
	currentDamage.SetString("1") // This needs to be calculated after upgrade info retrieval

	maxDamage := new(big.Float)
	maxDamage.SetString(gameData.MaxDamage)

	currentLevel := gameData.CurrentLevel

	maxLevel := gameData.MaxLevel

	currentStage := gameData.CurrentStage

	maxStage := gameData.MaxStage

	planetsDestroyed := new(big.Float)
	planetsDestroyed.SetString(gameData.PlanetsDestroyed)

	planet := Planet{
		Name:          "Planet_name",
		CurrentHealth: new(big.Float),
		MaxHealth:     new(big.Float),
		Gold:          new(big.Float),
	}

	store := map[string]StoreUpgrade{
		"1": {
			Level:  gameData.Store["1"].Level,
			Cost:   new(big.Float),
			Base:   new(big.Float),
			Damage: new(big.Float),
		},
		"2": {
			Level:  gameData.Store["2"].Level,
			Cost:   new(big.Float),
			Base:   new(big.Float),
			Damage: new(big.Float),
		},
		"3": {
			Level:  gameData.Store["3"].Level,
			Cost:   new(big.Float),
			Base:   new(big.Float),
			Damage: new(big.Float),
		},
		"4": {
			Level:  gameData.Store["4"].Level,
			Cost:   new(big.Float),
			Base:   new(big.Float),
			Damage: new(big.Float),
		},
	}

	store["1"].Base.SetString("10")
	store["2"].Base.SetString("1000")
	store["3"].Base.SetString("1.0e+5")
	store["4"].Base.SetString("1.0e+8")

	ship := map[string]ShipUpgrade{
		"1": {
			Level:      gameData.Ship["1"].Level,
			Cost:       new(big.Float),
			Multiplier: 1.0,
			Base:       new(big.Float),
			Damage:     new(big.Float),
		},
		"2": {
			Level:      gameData.Ship["2"].Level,
			Cost:       new(big.Float),
			Multiplier: 1.0,
			Base:       new(big.Float),
			Damage:     new(big.Float),
		},
		"3": {
			Level:      gameData.Ship["3"].Level,
			Cost:       new(big.Float),
			Multiplier: 1.0,
			Base:       new(big.Float),
			Damage:     new(big.Float),
		},
		"4": {
			Level:      gameData.Ship["4"].Level,
			Cost:       new(big.Float),
			Multiplier: 1.0,
			Base:       new(big.Float),
			Damage:     new(big.Float),
		},
	}

	// These bases costs are later to be changed
	// Depending on what function will the upgrade provide
	ship["1"].Base.SetString("10")
	ship["2"].Base.SetString("100")
	ship["3"].Base.SetString("1000")
	ship["4"].Base.SetString("10000")

	return &Game{
		Id:               id,
		Gold:             gold,
		Diamonds:         diamonds,
		CurrentDamage:    currentDamage,
		MaxDamage:        maxDamage,
		CurrentLevel:     currentLevel,
		MaxLevel:         maxLevel,
		CurrentStage:     currentStage,
		MaxStage:         maxStage,
		PlanetsDestroyed: planetsDestroyed,
		Planet:           planet,
		Store:            store,
		Ship:             ship,
		Ch:               make(chan bool),
	}
}
