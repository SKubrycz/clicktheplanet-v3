package main

import (
	"math/big"
)

type ShipUpgrade struct {
	Level      *big.Float
	Multiplier float64
	Damage     *big.Float
}

type StoreUpgrade struct {
	Level  *big.Float
	Damage *big.Float
}

type Planet struct {
	Name   string
	Health big.Float
}

type Game struct {
	Gold             big.Float
	Diamonds         int
	CurrentDamage    big.Float
	MaxDamage        big.Float
	CurrentLevel     int
	MaxLevel         int
	CurrentStage     int
	MaxStage         int
	PlanetsDestroyed big.Float // rounded to 0 decimal places
	Store            map[string]StoreUpgrade
	Ship             map[string]ShipUpgrade
}

func NewGame() *Game {
	// ... here data from database

	// big numbers initialization and db data assignment to variables
	gold := new(big.Float)
	var diamonds int
	currentDamage := new(big.Float)
	maxDamage := new(big.Float)
	var currentLevel int
	var maxLevel int
	var currentStage int
	var maxStage int
	planetsDestroyed := new(big.Float)
	store := map[string]StoreUpgrade{
		"1": StoreUpgrade{
			Level:  new(big.Float),
			Damage: new(big.Float),
		},
		"2": StoreUpgrade{
			Level:  new(big.Float),
			Damage: new(big.Float),
		},
		"3": StoreUpgrade{
			Level:  new(big.Float),
			Damage: new(big.Float),
		},
		"4": StoreUpgrade{
			Level:  new(big.Float),
			Damage: new(big.Float),
		},
	}
	ship := map[string]ShipUpgrade{
		"1": ShipUpgrade{
			Level:      new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
		},
		"2": ShipUpgrade{
			Level:      new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
		},
		"3": ShipUpgrade{
			Level:      new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
		},
		"4": ShipUpgrade{
			Level:      new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
		},
	}

	return &Game{
		Gold:             *gold,
		Diamonds:         diamonds,
		CurrentDamage:    *currentDamage,
		MaxDamage:        *maxDamage,
		CurrentLevel:     currentLevel,
		MaxLevel:         maxLevel,
		CurrentStage:     currentStage,
		MaxStage:         maxStage,
		PlanetsDestroyed: *planetsDestroyed,
		Store:            store,
		Ship:             ship,
	}
}
