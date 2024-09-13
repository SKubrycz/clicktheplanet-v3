package main

import (
	"math/big"
)

type ShipUpgrade struct {
	Level      big.Int
	Multiplier float64
	Damage     big.Float
}

type Ship struct {
	Upgrades map[string]ShipUpgrade
}

type StoreUpgrade struct {
	Level  big.Int
	Damage big.Float
}

type Store struct {
	Upgrades map[string]StoreUpgrade
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
	PlanetsDestroyed big.Int
	Store            Store
	Ship             Ship
}
