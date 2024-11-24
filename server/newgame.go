package main

import (
	"fmt"
	"math/big"
)

func NewGame(gameData *GameData) *Game {
	fmt.Println("ASSIGNING gameData")

	id := gameData.Id

	gold := new(big.Float)
	gold.SetString(gameData.Gold)

	diamonds := gameData.Diamonds

	currentDamage := new(big.Float)
	currentDamage.SetString("1")

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
		DiamondPlanet: DiamondPlanet{
			Diamonds:        1,
			Chance:          0.001,
			IsDiamondPlanet: false,
		},
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
		16: {
			Level:      gameData.Store[16].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		17: {
			Level:      gameData.Store[17].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		18: {
			Level:      gameData.Store[18].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		19: {
			Level:      gameData.Store[19].Level,
			Cost:       new(big.Float),
			BaseCost:   new(big.Float),
			Damage:     new(big.Float),
			BaseDamage: new(big.Float),
			Locked:     true,
		},
		20: {
			Level:      gameData.Store[20].Level,
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
	store[16].BaseCost.SetString("1.0e+43")
	store[17].BaseCost.SetString("1.0e+50")
	store[18].BaseCost.SetString("1.0e+59")
	store[19].BaseCost.SetString("1.0e+69")
	store[20].BaseCost.SetString("1.0e+80")

	for k := range len(store) {
		store[k+1].Cost.SetString(store[k+1].BaseCost.String())
	}

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
	store[16].BaseDamage.SetString("5.0e+28")
	store[17].BaseDamage.SetString("1.0e+31")
	store[18].BaseDamage.SetString("5.0e+33")
	store[19].BaseDamage.SetString("1.0e+37")
	store[20].BaseDamage.SetString("2.0e+40")

	for k := range len(store) {
		store[k+1].Damage.SetString(store[k+1].BaseDamage.String())
	}

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
