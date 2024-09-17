package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ShipUpgrade struct {
	Level      *big.Float
	Multiplier float64
	Base       int
	Damage     *big.Float
}

type StoreUpgrade struct {
	Level  *big.Float
	Base   int
	Damage *big.Float
}

type Planet struct {
	Name          string
	CurrentHealth *big.Float
	MaxHealth     *big.Float
	Gold          *big.Float
}

type Game struct {
	Gold             *big.Float
	Diamonds         int
	CurrentDamage    *big.Float
	MaxDamage        *big.Float
	CurrentLevel     int
	MaxLevel         int
	CurrentStage     int
	MaxStage         int
	PlanetsDestroyed *big.Float // rounded to 0 decimal places
	Planet           Planet
	Store            map[string]StoreUpgrade
	Ship             map[string]ShipUpgrade
}

type ShipUpgradeData struct {
	Level      string
	Multiplier float64
	Damage     string
}

type StoreUpgradeData struct {
	Level  string
	Damage string
}

// For data retrieved from the database
type GameData struct {
	Gold             string
	Diamonds         int
	MaxDamage        string
	CurrentLevel     int
	MaxLevel         int
	CurrentStage     int
	MaxStage         int
	PlanetsDestroyed string
	Store            map[string]StoreUpgradeData
	Ship             map[string]ShipUpgradeData
}

func NewGame(gameData *GameData) *Game {
	// big numbers initialization and db data assignment to variables
	gold := new(big.Float)
	gold.SetString(gameData.Gold)

	var diamonds int
	diamonds = gameData.Diamonds

	currentDamage := new(big.Float)
	currentDamage.SetString("1") // This needs to be calculated after upgrade info retrieval

	maxDamage := new(big.Float)
	maxDamage.SetString(gameData.MaxDamage)

	var currentLevel int
	currentLevel = gameData.CurrentLevel

	var maxLevel int
	maxLevel = gameData.MaxLevel

	var currentStage int
	currentStage = gameData.CurrentStage

	var maxStage int
	maxStage = gameData.MaxStage

	planetsDestroyed := new(big.Float)
	planetsDestroyed.SetString(gameData.PlanetsDestroyed)

	planet := Planet{
		Name:          "Planet_name",
		CurrentHealth: new(big.Float),
		MaxHealth:     new(big.Float),
		Gold:          new(big.Float),
	}

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
	store["1"].Level.SetString(gameData.Store["1"].Level)
	store["2"].Level.SetString(gameData.Store["2"].Level)
	store["3"].Level.SetString(gameData.Store["3"].Level)
	store["4"].Level.SetString(gameData.Store["4"].Level)
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
	ship["1"].Level.SetString(gameData.Ship["1"].Level)
	ship["2"].Level.SetString(gameData.Ship["2"].Level)
	ship["3"].Level.SetString(gameData.Ship["3"].Level)
	ship["4"].Level.SetString(gameData.Ship["4"].Level)

	return &Game{
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
	}
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
