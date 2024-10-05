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

	damageDone := DamageDone{
		Damage:   big.NewFloat(0),
		Critical: false,
	}

	dps := new(big.Float)
	dps.SetString("1")

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

	store := map[int]StoreUpgrade{
		1: {
			Level:      gameData.Store[1].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
		2: {
			Level:      gameData.Store[2].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
		3: {
			Level:      gameData.Store[3].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
		4: {
			Level:      gameData.Store[4].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
	}

	store[1].BaseCost.SetString("10")
	store[2].BaseCost.SetString("500")
	store[3].BaseCost.SetString("10000")
	store[4].BaseCost.SetString("1000000")
	store[1].Cost.SetString("10")
	store[2].Cost.SetString("500")
	store[3].Cost.SetString("10000")
	store[4].Cost.SetString("1000000")

	store[1].BaseDamage.SetString("1")
	store[2].BaseDamage.SetString("100")
	store[3].BaseDamage.SetString("5000")
	store[4].BaseDamage.SetString("50000")
	store[1].Damage.SetString("1")
	store[2].Damage.SetString("100")
	store[3].Damage.SetString("5000")
	store[4].Damage.SetString("50000")

	ship := map[int]ShipUpgrade{
		1: {
			Level:      gameData.Ship[1].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 0.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
		2: {
			Level:      gameData.Ship[2].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
		3: {
			Level:      gameData.Ship[3].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 0.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
		4: {
			Level:      gameData.Ship[4].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
		},
	}

	// These bases costs are later to be changed
	// Depending on what function will the upgrade provide
	ship[1].BaseCost.SetString("10")
	ship[2].BaseCost.SetString("100")
	ship[3].BaseCost.SetString("1000")
	ship[4].BaseCost.SetString("10000")
	ship[1].Cost.SetString("10")
	ship[2].Cost.SetString("100")
	ship[3].Cost.SetString("1000")
	ship[4].Cost.SetString("10000")

	ship[1].BaseDamage.SetString("10")
	ship[2].BaseDamage.SetString("50")
	ship[3].BaseDamage.SetString("500")
	ship[4].BaseDamage.SetString("10000")
	ship[1].Damage.SetString("10")
	ship[2].Damage.SetString("50")
	ship[3].Damage.SetString("500")
	ship[4].Damage.SetString("10000")

	return &Game{
		Id:               id,
		Gold:             gold,
		Diamonds:         diamonds,
		CurrentDamage:    currentDamage,
		MaxDamage:        maxDamage,
		DamageDone:       damageDone,
		Dps:              dps,
		CurrentLevel:     currentLevel,
		MaxLevel:         maxLevel,
		CurrentStage:     currentStage,
		MaxStage:         maxStage,
		PlanetsDestroyed: planetsDestroyed,
		Planet:           planet,
		Store:            store,
		Ship:             ship,
		Ch:               make(chan string),
	}
}
