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

	diamonds := new(big.Float)
	diamonds.SetString(gameData.Diamonds)

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
		Diamonds:      new(big.Float),
		IsBoss:        false,
		DiamondPlanet: DiamondPlanet{
			Diamonds:        1,
			BaseChance:      0.001,
			Chance:          0.001,
			IsDiamondPlanet: false,
		},
	}

	store := map[int]StoreUpgrade{}

	for k := range len(gameData.Store) {
		s := new(StoreUpgrade)
		s.Level = gameData.Store[k+1].Level
		s.Cost = new(big.Float)
		s.BaseCost = new(big.Float)
		s.Damage = new(big.Float)
		s.BaseDamage = new(big.Float)
		s.Locked = true

		store[k+1] = *s
	}

	store[1].BaseCost.SetString("10")
	store[2].BaseCost.SetString("200")
	store[3].BaseCost.SetString("10000")
	store[4].BaseCost.SetString("250000")
	store[5].BaseCost.SetString("5.0e+7")
	store[6].BaseCost.SetString("7.5e+10")
	store[7].BaseCost.SetString("2.5e+13")
	store[8].BaseCost.SetString("2.0e+19")
	store[9].BaseCost.SetString("1.0e+26")
	store[10].BaseCost.SetString("1.0e+34")
	store[11].BaseCost.SetString("1.0e+43")
	store[12].BaseCost.SetString("1.0e+53")
	store[13].BaseCost.SetString("1.0e+64")
	store[14].BaseCost.SetString("1.0e+75")
	store[15].BaseCost.SetString("1.0e+87")
	store[16].BaseCost.SetString("1.0e+100")
	store[17].BaseCost.SetString("1.0e+114")
	store[18].BaseCost.SetString("1.0e+130")
	store[19].BaseCost.SetString("1.0e+148")
	store[20].BaseCost.SetString("1.0e+168")
	store[21].BaseCost.SetString("1.0e+190")
	store[22].BaseCost.SetString("1.0e+214")
	store[23].BaseCost.SetString("1.0e+240")
	store[24].BaseCost.SetString("1.0e+290")
	store[25].BaseCost.SetString("1.0e+360")
	store[26].BaseCost.SetString("1.0e+450")
	store[27].BaseCost.SetString("1.0e+560")
	store[28].BaseCost.SetString("1.0e+690")
	store[29].BaseCost.SetString("1.0e+840")
	store[30].BaseCost.SetString("1.0e+1010")

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
	store[8].BaseDamage.SetString("1.0e+13")
	store[9].BaseDamage.SetString("1.0e+16")
	store[10].BaseDamage.SetString("1.0e+20")
	store[11].BaseDamage.SetString("1.0e+24")
	store[12].BaseDamage.SetString("1.0e+29")
	store[13].BaseDamage.SetString("2.5e+34")
	store[14].BaseDamage.SetString("1.0e+40")
	store[15].BaseDamage.SetString("2.0e+46")
	store[16].BaseDamage.SetString("5.0e+52")
	store[17].BaseDamage.SetString("1.0e+59")
	store[18].BaseDamage.SetString("5.0e+66")
	store[19].BaseDamage.SetString("1.0e+74")
	store[20].BaseDamage.SetString("2.0e+82")
	store[21].BaseDamage.SetString("1.0e+91")
	store[22].BaseDamage.SetString("1.0e+100")
	store[23].BaseDamage.SetString("2.5e+110")
	store[24].BaseDamage.SetString("1.0e+125")
	store[25].BaseDamage.SetString("1.0e+145")
	store[26].BaseDamage.SetString("1.0e+170")
	store[27].BaseDamage.SetString("1.0e+200")
	store[28].BaseDamage.SetString("1.0e+235")
	store[29].BaseDamage.SetString("1.0e+275")
	store[30].BaseDamage.SetString("1.0e+320")

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
			Cost:       big.NewFloat(1),
			Constant:   1.5,
		},
		2: {
			Level:      gameData.DiamondUpgrade[2].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       big.NewFloat(1),
			Constant:   1.5,
		},
		3: {
			Level:      gameData.DiamondUpgrade[3].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       big.NewFloat(1),
			Constant:   1.5,
		},
		4: {
			Level:      gameData.DiamondUpgrade[4].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       big.NewFloat(1),
			Constant:   1.5,
		},
		5: {
			Level:      gameData.DiamondUpgrade[5].Level,
			Multiplier: big.NewFloat(1.0),
			BaseCost:   1,
			Cost:       big.NewFloat(1),
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
