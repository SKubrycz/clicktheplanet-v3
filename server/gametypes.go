package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ShipUpgrade struct {
	Level      int64
	Cost       *big.Float
	BaseCost   *big.Float
	Multiplier float64
	Damage     *big.Float
	BaseDamage *big.Float
}

type StoreUpgrade struct {
	Level      int64
	Cost       *big.Float
	BaseCost   *big.Float
	Damage     *big.Float
	BaseDamage *big.Float
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
	Dps              *big.Float
	CurrentLevel     int64
	MaxLevel         int64
	CurrentStage     uint8
	MaxStage         uint8
	PlanetsDestroyed *big.Float
	Planet           Planet
	Store            map[int]StoreUpgrade
	Ship             map[int]ShipUpgrade
	Ch               chan string
}

func (g *Game) ClickThePlanet(dmg *big.Float) {
	g.Planet.CurrentHealth.Sub(g.Planet.CurrentHealth, dmg)
	x := big.NewFloat(0)
	if g.Planet.CurrentHealth.Cmp(x) <= 0 {
		g.Advance()
		g.CalculatePlanetHealth()
		g.AddPlanetDestroyed()
		g.CalculateGoldEarned()
		g.AddCurrentGold()
		g.Ch <- "click"
	}
}

func (g *Game) CalculateCurrentDamage() {
	res := big.NewFloat(1) //base dmg
	for k := range g.Store {
		if g.Store[k].Level > 0 {
			res.Add(res, g.Store[k].Damage)
		}
	}
	g.CurrentDamage = res
	if g.CurrentDamage.Cmp(g.MaxDamage) > 0 {
		g.MaxDamage = g.CurrentDamage
	}

	g.Dps.Mul(g.Ship[1].Damage, big.NewFloat(10.0)) // per second
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
	return percentage
}

func (g *Game) CalculatePlanetHealth() {
	f := 1.3

	exp := float64(g.CurrentLevel - 1)
	pow := math.Pow(f, exp)

	result := new(big.Float).SetFloat64(pow)
	result.Mul(result, big.NewFloat(10))

	g.ConvertNumber(result, g.Planet.MaxHealth)
	g.ConvertNumber(result, g.Planet.CurrentHealth)
}

func (g *Game) GeneratePlanetName() {
	lns := g.CurrentLevel + int64(g.CurrentStage)

	rand := (123*lns + 123768173) % 7847681738
	stringLength := int((rand % 5) + 4)

	var byteName bytes.Buffer
	// 48 - 57 range 9
	// 65 - 90 range 25
	// (ASCII) ^ inclusive

	randNumber := (5634*lns + 5609872012) % 5609872012923
	randLetter := (8946*lns + 5609872012) % 5609872012923

	for i := 0; i < stringLength; i++ {
		randNumber = (2119*randNumber + 5609872012) % 5609872012923
		randLetter = (8946*randLetter + 1238971239) % 5609872012923

		number := (randNumber % 9) + 48
		letter := (randLetter % 25) + 65

		if randNumber > randLetter {
			byteName.WriteByte(byte(number))
		} else if randNumber <= randLetter {
			byteName.WriteByte(byte(letter))
		}
	}

	g.Planet.Name = byteName.String()
}

func (g *Game) UpgradeStore(index int, levels int) string {
	if g.Gold.Cmp(g.Store[index].Cost) >= 0 {
		f := 1.03
		bulkCost := new(big.Float)

		for i := 1; i <= levels; i++ {
			pow := math.Pow(f, float64(g.Store[index].Level))
			bigPow := big.NewFloat(pow)
			cost := new(big.Float)

			cost.Mul(g.Store[index].BaseCost, bigPow)

			bulkCost.Add(bulkCost, cost)
		}
		fmt.Println("   ---> bulkCost: ", bulkCost)

		if bulkCost.Cmp(g.Gold) > 0 {
			return "insufficient resources"
		}

		fmt.Printf("\n %v \n", bulkCost)

		if entry, ok := g.Store[index]; ok {
			entry.Level += int64(levels)
			g.Store[index] = entry
			result := new(big.Float)
			result.Sub(g.Gold, bulkCost) // g.Store[index].Cost * levels
			g.ConvertNumber(result, g.Gold)
		}
		g.Ch <- "upgrade"
	} else {
		return "insufficient resources"
	}
	return ""
}

func (g *Game) CalculateStore(index int) {
	// If the index != -1, calculate the specific StoreUpgrade
	// else just calculate the whole map
	if index > len(g.Store) {
		return
	}
	// constant for the formula
	f := 1.03

	if index != -1 {
		if g.Store[index].Level > 0 {
			// cost = baseCost * 1.03 ** (level - 1)
			pow := math.Pow(f, float64(g.Store[index].Level))
			bigPow := big.NewFloat(pow)
			cost := new(big.Float)

			cost.Mul(g.Store[index].BaseCost, bigPow)

			g.ConvertNumber(cost, g.Store[index].Cost)

			//damage = baseDmg * level
			bigLevel := big.NewFloat(float64(g.Store[index].Level))
			g.Store[index].Damage.Mul(g.Store[index].BaseDamage, bigLevel)
		}
	} else if index == -1 {
		for k := range g.Store {
			if g.Store[k].Level > 0 {
				pow := math.Pow(f, float64(g.Store[k].Level))
				bigPow := big.NewFloat(pow)

				cost := new(big.Float)
				cost.Mul(g.Store[k].BaseCost, bigPow)

				g.ConvertNumber(cost, g.Store[k].Cost)

				bigLevel := big.NewFloat(float64(g.Store[k].Level))
				g.Store[k].Damage.Mul(g.Store[k].BaseDamage, bigLevel)
			}
		}
	} else {
		return
	}

	g.CalculateShip(1)
	g.CalculateCurrentDamage()
}

func (g *Game) UpgradeShip(index int) string {
	if g.Gold.Cmp(g.Ship[index].Cost) >= 0 {
		if entry, ok := g.Ship[index]; ok {
			entry.Level += 1
			g.Ship[index] = entry
			result := new(big.Float)
			result.Sub(g.Gold, g.Ship[index].Cost)
			g.ConvertNumber(result, g.Gold)
		}
		g.Ch <- "upgrade"
	} else {
		return "insufficient resources"
	}
	g.CalculateCurrentDamage()

	return ""
}

func (g *Game) CalculateShip(index int) {
	if index > len(g.Ship) {
		return
	}

	f := 1.5

	if index == 1 {
		// cost = baseCost * 1.5 ** (level - 1)
		pow := math.Pow(f, float64(g.Ship[index].Level))
		bigPow := big.NewFloat(pow)

		cost := new(big.Float)
		cost.Mul(g.Ship[index].BaseCost, bigPow)
		g.ConvertNumber(cost, g.Ship[index].Cost)

		if entry, ok := g.Ship[index]; ok {
			// every 100ms
			entry.Multiplier = toFixed((0.001 * float64(g.Ship[index].Level)), 3) // 10
			g.Ship[index] = entry
		}

		//damage = currentDamage * multiplier
		bigMultiplier := big.NewFloat(g.Ship[index].Multiplier)
		g.Ship[index].Damage.Mul(g.CurrentDamage, bigMultiplier) // per 100ms
	} else if index != -1 {
		// Level      int64
		// Cost       *big.Float
		// BaseCost   *big.Float
		// Multiplier float64
		// Damage     *big.Float
		// BaseDamage *big.Float

		// cost = baseCost * 1.5 ** (level - 1)
		pow := math.Pow(f, float64(g.Ship[index].Level))
		bigPow := big.NewFloat(pow)

		cost := new(big.Float)
		cost.Mul(g.Ship[index].BaseCost, bigPow)
		g.ConvertNumber(cost, g.Ship[index].Cost)

		if entry, ok := g.Ship[index]; ok {
			// every 100ms
			entry.Multiplier = toFixed((0.001 * float64(g.Ship[index].Level)), 3) // 10
			g.Ship[index] = entry
		}

		//damage = currentDamage * multiplier
		bigMultiplier := big.NewFloat(g.Ship[index].Multiplier)
		g.Ship[index].Damage.Mul(g.CurrentDamage, bigMultiplier) // per 100ms

	} else if index == -1 {
		for k := range g.Ship {
			// cost = baseCost * 1.5 ** (level - 1)
			pow := math.Pow(f, float64(g.Ship[k].Level))
			bigPow := big.NewFloat(pow)

			cost := new(big.Float)
			cost.Mul(g.Ship[k].BaseCost, bigPow)
			g.ConvertNumber(cost, g.Ship[k].Cost)

			if entry, ok := g.Ship[k]; ok {
				// every 100ms
				entry.Multiplier = toFixed((0.001 * float64(g.Ship[k].Level)), 3) // 10
				g.Ship[k] = entry
			}

			//damage = currentDamage * multiplier
			bigMultiplier := big.NewFloat(g.Ship[k].Multiplier)
			g.Ship[k].Damage.Mul(g.CurrentDamage, bigMultiplier) // per 100ms
		}
	} else {
		return
	}

	g.CalculateCurrentDamage()

	// fmt.Printf("\n %v <- currentDamage inside ship \n", g.CurrentDamage)
	// fmt.Printf("\nship1damage: %v\n", g.Ship[1].Damage)
	// fmt.Printf("\nDPS: %v\n", g.Dps)
}

func (g *Game) Advance() {
	g.CurrentStage++
	if g.CurrentStage > 10 {
		g.CurrentLevel++
		g.CurrentStage = 1
	}
	if g.CurrentLevel > g.MaxLevel {
		g.MaxLevel = g.CurrentLevel
	}
	if g.CurrentLevel == g.MaxLevel {
		g.MaxStage = g.CurrentStage
	}
	g.GeneratePlanetName()
}

func (g *Game) PreviousLevel() {
	// g.GeneratePlanetName()
}

func (g *Game) NextLevel() {
	// g.GeneratePlanetName()
}

func (g *Game) CalculateGoldEarned() {
	f := 1.4

	exp := float64(g.CurrentLevel)
	pow := math.Pow(f, exp)

	result := new(big.Float).SetFloat64(pow)
	result.Mul(result, big.NewFloat(10))

	g.ConvertNumber(result, g.Planet.Gold)
}

func (g *Game) AddCurrentGold() {
	result := new(big.Float)
	result.Add(g.Gold, g.Planet.Gold)

	g.ConvertNumber(result, g.Gold)
}

func (g *Game) AddPlanetDestroyed() {
	add := new(big.Float)
	one := new(big.Float)
	one.SetString("1")
	add.Add(g.PlanetsDestroyed, one)
	g.PlanetsDestroyed = add
}

func (g *Game) ConvertNumber(res *big.Float, dest *big.Float) {
	intRes, _ := res.Int(nil)

	if len(intRes.String()) > 6 {
		dest.SetString(res.Text('e', 3))
	} else {
		dest.SetString(res.Text('f', 0))
	}
}

// DisplayNumber returns a big number string to be sent to client
func (g *Game) DisplayNumber(num *big.Float) string {
	intNum, _ := num.Int(nil)

	if len(intNum.String()) > 6 {
		return num.Text('e', 3)
	} else {
		return num.Text('f', 0)
	}
}

func toFixed(num float64, prec int) float64 {
	output := math.Pow(10, float64(prec))
	round := int((num * output) + math.Copysign(0.5, num))
	return float64(round) / output
}
