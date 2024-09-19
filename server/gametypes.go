package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ShipUpgrade struct {
	Level      int64
	Cost       *big.Float
	Multiplier float64
	Base       *big.Float
	Damage     *big.Float
}

type StoreUpgrade struct {
	Level  int64
	Cost   *big.Float
	Base   *big.Float
	Damage *big.Float
}

type Planet struct {
	Name          string
	CurrentHealth *big.Float
	MaxHealth     *big.Float
	Gold          *big.Float
}

type Game struct {
	Id               int64
	Gold             *big.Float
	Diamonds         int64
	CurrentDamage    *big.Float
	MaxDamage        *big.Float
	CurrentLevel     int64
	MaxLevel         int64
	CurrentStage     uint8
	MaxStage         uint8
	PlanetsDestroyed *big.Float // rounded to 0 decimal places
	Planet           Planet
	Store            map[string]StoreUpgrade
	Ship             map[string]ShipUpgrade
	Ch               chan bool
}

func (g *Game) ClickThePlanet() {
	g.Planet.CurrentHealth.Sub(g.Planet.CurrentHealth, g.CurrentDamage)
	x := big.NewFloat(0)
	if g.Planet.CurrentHealth.Cmp(x) <= 0 {
		g.Advance()
		g.CalculatePlanetHealth()
		g.AddPlanetDestroyed()
		g.CalculateGoldEarned()
		g.AddCurrentGold()
		g.Ch <- true
	}
}

func (g *Game) GetHealthPercent() int {
	hundred := big.NewFloat(100)
	healthMultiplied := new(big.Float)
	healthMultiplied.SetString(g.Planet.CurrentHealth.Text('f', 3))
	healthMultiplied.Mul(healthMultiplied, hundred)
	healthMultiplied.Quo(healthMultiplied, g.Planet.MaxHealth)
	intHealth, _ := healthMultiplied.Int(nil)
	intText := intHealth.Text(10)
	percentage, _ := strconv.Atoi(intText)
	fmt.Println(percentage)
	return percentage
}

func (g *Game) CalculatePlanetHealth() {
	fmt.Println("CalculatePlanetHealth")

	// Formula
	f := 1.3

	exp := float64(g.CurrentLevel - 1)
	pow := math.Pow(f, exp)

	result := new(big.Float).SetFloat64(pow)
	result.Mul(result, big.NewFloat(10))

	intResult, _ := result.Int(nil)

	if len(intResult.String()) > 6 {
		fmt.Println(result.Text('e', 3))
		g.Planet.MaxHealth.SetString(result.Text('e', 3))
		g.Planet.CurrentHealth.SetString(result.Text('e', 3)) // set planet to 100%hp
	} else {
		fmt.Println(result.Text('f', 0))
		g.Planet.MaxHealth.SetString(result.Text('f', 0))
		g.Planet.CurrentHealth.SetString(result.Text('f', 0)) // set planet to 100%hp
	}
}

func (g *Game) CalculateStore(index int) {
	// If the index != -1, calculate the specific StoreUpgrade
	// else just calculate the whole map
	indexStr := strconv.Itoa(index)
	if index > len(g.Store) {
		return
	}
	if index != -1 {
		if g.Store[indexStr].Level > 0 {
			// cost = base * 1.03 ** (level - 1)
			// g.Store[indexStr].Cost
			// g.Store[indexStr].Base
			//f := 1.03
		}
	} else if index == -1 {
		for k := range g.Store {
			if g.Store[k].Level > 0 {

			}
		}
	} else {
		return
	}
}

func (g *Game) Advance() {
	// return new Planet.Health; new Planet.Name
	// set g.CurrentStage to 1; if g.CurrentLevel > MaxLevel then g.MaxStage = g.CurrentStage
	// calculate max health and current health equal to max health
	g.CurrentStage++
	if g.CurrentStage > 10 {
		g.CurrentLevel++ // for each level up / stage up, save progress to database
		g.CurrentStage = 1
	}
	if g.CurrentLevel > g.MaxLevel {
		g.MaxLevel = g.CurrentLevel
	}
	if g.CurrentLevel == g.MaxLevel {
		if g.CurrentStage > g.MaxStage {
			g.MaxStage = g.CurrentStage
		}
	}

}

func (g *Game) PreviousLevel() {

}

func (g *Game) NextLevel() {

}

func (g *Game) Upgrade( /* store/ship */ ) {

}

func (g *Game) CalculateGoldEarned() {
	f := 1.05

	exp := float64(g.CurrentLevel)
	pow := math.Pow(f, exp)

	result := new(big.Float).SetFloat64(pow)
	result.Mul(result, big.NewFloat(10))

	intResult, _ := result.Int(nil)

	if len(intResult.String()) > 6 {
		fmt.Println(result.Text('e', 3))
		g.Planet.Gold.SetString(result.Text('e', 3))
	} else {
		fmt.Println(result.Text('f', 0))
		g.Planet.Gold.SetString(result.Text('f', 0))
	}
}

func (g *Game) AddCurrentGold() {
	result := new(big.Float)
	result.Add(g.Gold, g.Planet.Gold)

	g.Gold = result
}

func (g *Game) AddPlanetDestroyed() {
	add := new(big.Float)
	one := new(big.Float)
	one.SetString("1")
	add.Add(g.PlanetsDestroyed, one)
	g.PlanetsDestroyed = add
}
