package main

import (
	"fmt"
	"math/big"
)

func NewGame(gameData *GameData) *Game {
	fmt.Println("ASSIGNING gameData")

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
		Damage:   currentDamage,
		Critical: false,
	}

	dps := new(big.Float)
	dps.SetString("1")

	currentLevel := gameData.CurrentLevel

	maxLevel := gameData.MaxLevel

	currentStage := gameData.CurrentStage

	maxStage := gameData.MaxStage

	diamondUpgradesUnlocked := false

	planetsDestroyed := new(big.Float)
	planetsDestroyed.SetString(gameData.PlanetsDestroyed)

	planet := Planet{
		Name:          "Planet_name",
		CurrentHealth: new(big.Float),
		MaxHealth:     new(big.Float),
		Gold:          new(big.Float),
		IsBoss:        false,
	}

	store := map[int]StoreUpgrade{
		1: {
			Level:      gameData.Store[1].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		2: {
			Level:      gameData.Store[2].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		3: {
			Level:      gameData.Store[3].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		4: {
			Level:      gameData.Store[4].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		5: {
			Level:      gameData.Store[5].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		6: {
			Level:      gameData.Store[6].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		7: {
			Level:      gameData.Store[7].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		8: {
			Level:      gameData.Store[8].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		9: {
			Level:      gameData.Store[9].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		10: {
			Level:      gameData.Store[10].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		11: {
			Level:      gameData.Store[11].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		12: {
			Level:      gameData.Store[12].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		13: {
			Level:      gameData.Store[13].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		14: {
			Level:      gameData.Store[14].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		15: {
			Level:      gameData.Store[15].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
	}

	store[1].BaseCost.SetString("10")
	store[2].BaseCost.SetString("200")
	store[3].BaseCost.SetString("10000")
	store[4].BaseCost.SetString("250000")
	store[5].BaseCost.SetString("5.0e+7")
	store[6].BaseCost.SetString("7.5e+10")
	store[7].BaseCost.SetString("2.5e+13")
	store[8].BaseCost.SetString("2.0e+16")
	store[9].BaseCost.SetString("1.0e+19")
	store[10].BaseCost.SetString("1.0e+22")
	store[11].BaseCost.SetString("1.0e+25")
	store[12].BaseCost.SetString("1.0e+28")
	store[13].BaseCost.SetString("1.0e+31")
	store[14].BaseCost.SetString("1.0e+35")
	store[15].BaseCost.SetString("1.0e+39")

	store[1].Cost.SetString(store[1].BaseCost.String())
	store[2].Cost.SetString(store[2].BaseCost.String())
	store[3].Cost.SetString(store[3].BaseCost.String())
	store[4].Cost.SetString(store[4].BaseCost.String())
	store[5].Cost.SetString(store[5].BaseCost.String())
	store[6].Cost.SetString(store[6].BaseCost.String())
	store[7].Cost.SetString(store[7].BaseCost.String())
	store[8].Cost.SetString(store[8].BaseCost.String())
	store[9].Cost.SetString(store[9].BaseCost.String())
	store[10].Cost.SetString(store[10].BaseCost.String())
	store[11].Cost.SetString(store[11].BaseCost.String())
	store[12].Cost.SetString(store[12].BaseCost.String())
	store[13].Cost.SetString(store[13].BaseCost.String())
	store[14].Cost.SetString(store[14].BaseCost.String())
	store[15].Cost.SetString(store[15].BaseCost.String())

	store[1].BaseDamage.SetString("1")
	store[2].BaseDamage.SetString("25")
	store[3].BaseDamage.SetString("400")
	store[4].BaseDamage.SetString("10000")
	store[5].BaseDamage.SetString("1.0e+6")
	store[6].BaseDamage.SetString("1.0e+8")
	store[7].BaseDamage.SetString("1.0e+10")
	store[8].BaseDamage.SetString("1.0e+12")
	store[9].BaseDamage.SetString("1.0e+14")
	store[10].BaseDamage.SetString("1.0e+16")
	store[11].BaseDamage.SetString("1.0e+18")
	store[12].BaseDamage.SetString("1.0e+20")
	store[13].BaseDamage.SetString("1.0e+22")
	store[14].BaseDamage.SetString("1.0e+24")
	store[15].BaseDamage.SetString("2.0e+26")

	store[1].Damage.SetString(store[1].BaseDamage.String())
	store[2].Damage.SetString(store[2].BaseDamage.String())
	store[3].Damage.SetString(store[3].BaseDamage.String())
	store[4].Damage.SetString(store[4].BaseDamage.String())
	store[5].Damage.SetString(store[5].BaseDamage.String())
	store[6].Damage.SetString(store[6].BaseDamage.String())
	store[7].Damage.SetString(store[7].BaseDamage.String())
	store[8].Damage.SetString(store[8].BaseDamage.String())
	store[9].Damage.SetString(store[9].BaseDamage.String())
	store[10].Damage.SetString(store[10].BaseDamage.String())
	store[11].Damage.SetString(store[11].BaseDamage.String())
	store[12].Damage.SetString(store[12].BaseDamage.String())
	store[13].Damage.SetString(store[13].BaseDamage.String())
	store[14].Damage.SetString(store[14].BaseDamage.String())
	store[15].Damage.SetString(store[15].BaseDamage.String())

	// Constant field is needed to calculate cost
	ship := map[int]ShipUpgrade{
		1: {
			Level:      gameData.Ship[1].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 0.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
			Constant:   1.3,
		},
		2: {
			Level:      gameData.Ship[2].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
			Constant:   1.2,
		},
		3: {
			Level:      gameData.Ship[3].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 0.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
			Constant:   1.5,
		},
		4: {
			Level:      gameData.Ship[4].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Multiplier: 1.0,
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
			Constant:   1.1,
		},
	}

	// These bases costs are later to be changed
	// Depending on what function will the upgrade provide
	ship[1].BaseCost.SetString("100")
	ship[2].BaseCost.SetString("500")
	ship[3].BaseCost.SetString("10000")
	ship[4].BaseCost.SetString("100000")
	ship[1].Cost.SetString(ship[1].BaseCost.String())
	ship[2].Cost.SetString(ship[2].BaseCost.String())
	ship[3].Cost.SetString(ship[3].BaseCost.String())
	ship[4].Cost.SetString(ship[4].BaseCost.String())

	ship[1].BaseDamage.SetString("10")
	ship[2].BaseDamage.SetString("50")
	ship[3].BaseDamage.SetString("500")
	ship[4].BaseDamage.SetString("10000")
	ship[1].Damage.SetString(ship[1].BaseDamage.String())
	ship[2].Damage.SetString(ship[2].BaseDamage.String())
	ship[3].Damage.SetString(ship[3].BaseDamage.String())
	ship[4].Damage.SetString(ship[4].BaseDamage.String())

	diamondUpgrade := map[int]DiamondUpgrade{
		1: {
			Level:      gameData.DiamondUpgrade[1].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       1,
			Constant:   1.5,
		},
		2: {
			Level:      gameData.DiamondUpgrade[2].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       1,
			Constant:   1.5,
		},
		3: {
			Level:      gameData.DiamondUpgrade[3].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       1,
			Constant:   1.5,
		},
		4: {
			Level:      gameData.DiamondUpgrade[4].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       1,
			Constant:   1.5,
		},
		5: {
			Level:      gameData.DiamondUpgrade[5].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       1,
			Constant:   1.5,
		},
	}

	return &Game{
		Id:                      id,
		Gold:                    gold,
		Diamonds:                diamonds,
		CurrentDamage:           currentDamage,
		MaxDamage:               maxDamage,
		DamageDone:              damageDone,
		Dps:                     dps,
		CurrentLevel:            currentLevel,
		MaxLevel:                maxLevel,
		CurrentStage:            currentStage,
		MaxStage:                maxStage,
		DiamondUpgradesUnlocked: diamondUpgradesUnlocked,
		PlanetsDestroyed:        planetsDestroyed,
		Planet:                  planet,
		Store:                   store,
		Ship:                    ship,
		DiamondUpgrade:          diamondUpgrade,
		Ch:                      make(chan string),
	}
}
