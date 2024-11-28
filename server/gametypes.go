package main

import (
	"bytes"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
)

type DiamondUpgrade struct {
	Level      int64
	Multiplier *big.Float // float64
	BaseCost   int64
	Cost       *big.Float
	Constant   float64
}

type ShipUpgrade struct {
	Level          int64
	Cost           *big.Float
	BaseCost       *big.Float
	Multiplier     float64
	Damage         *big.Float
	BaseDamage     *big.Float
	Locked         bool
	Constant       float64 // for Cost formula
	DiamondUpgrade DiamondUpgrade
}

type StoreUpgrade struct {
	Level      int64
	Cost       *big.Float
	BaseCost   *big.Float
	Damage     *big.Float
	BaseDamage *big.Float
	Locked     bool
}

type DiamondPlanet struct {
	Diamonds        int64
	BaseChance      float64
	Chance          float64
	IsDiamondPlanet bool
}

type Planet struct {
	Name          string
	CurrentHealth *big.Float
	MaxHealth     *big.Float
	Gold          *big.Float
	Diamonds      *big.Float
	IsBoss        bool
	DiamondPlanet DiamondPlanet
}

type DamageDone struct {
	Damage   *big.Float
	Critical bool
}

type Game struct {
	Id                      int64
	Gold                    *big.Float
	Diamonds                *big.Float
	CurrentDamage           *big.Float
	MaxDamage               *big.Float
	DamageDone              DamageDone
	Dps                     *big.Float
	CurrentLevel            int64
	MaxLevel                int64
	CurrentStage            uint8
	MaxStage                uint8
	DiamondUpgradesUnlocked bool
	PlanetsDestroyed        *big.Float
	Planet                  Planet
	Store                   map[int]StoreUpgrade
	Ship                    map[int]ShipUpgrade
	DiamondUpgrade          map[int]DiamondUpgrade
	Ch                      chan string
}

func (g *Game) ClickThePlanet(dmg *big.Float, isClick bool) {
	if isClick {
		random := rand.Float64()
		if g.Ship[3].Multiplier > 0.0 && random <= g.Ship[3].Multiplier {
			dmg = g.Ship[3].Damage
			g.DamageDone.Damage = dmg
			g.DamageDone.Critical = true
		} else {
			g.DamageDone.Damage = dmg
			g.DamageDone.Critical = false
		}
		g.Planet.CurrentHealth.Sub(g.Planet.CurrentHealth, dmg)
	} else {
		g.Planet.CurrentHealth.Sub(g.Planet.CurrentHealth, dmg)
		g.DamageDone.Damage = dmg
		g.DamageDone.Critical = false
	}
	x := big.NewFloat(0)
	if g.Planet.CurrentHealth.Cmp(x) <= 0 {
		g.CalculateDiamondsEarned()
		g.CalculateGoldEarned()
		g.AddCurrentGold()
		g.Advance()
		g.CalculatePlanetHealth()
		g.AddPlanetDestroyed()
		g.CheckDiamondUpgradeUnlock()
		g.CheckStoreLock(-1)
		g.CheckShipLock(-1)
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
	bigMultiplier := big.NewFloat(g.Ship[2].Multiplier)
	res.Mul(res, bigMultiplier)
	if g.DiamondUpgrade[2].Level > 0 {
		res.Mul(res, g.DiamondUpgrade[2].Multiplier)
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
	f := 1.4

	exp := float64(g.CurrentLevel - 1)
	pow := math.Pow(f, exp)

	result := new(big.Float).SetFloat64(pow)
	if g.Planet.IsBoss {
		result.Mul(result, big.NewFloat(50))
	} else {
		result.Mul(result, big.NewFloat(10))
	}

	g.ConvertNumber(result, g.Planet.MaxHealth)
	g.ConvertNumber(result, g.Planet.CurrentHealth)
}

func (g *Game) GeneratePlanetName() {
	if !g.Planet.DiamondPlanet.IsDiamondPlanet {
		lns := (g.CurrentLevel * g.CurrentLevel) + int64(g.CurrentStage*g.CurrentStage)

		rand := (792*lns + 78916450) % 7847681738
		stringLength := int((rand % 5) + 4)

		var byteName bytes.Buffer
		// 48 - 57 range 9
		// 65 - 90 range 25
		// (ASCII) ^ inclusive

		randNumber := uint64((211986194*rand + 5609872012) % 5609872012923)
		randLetter := uint64((894676913*rand + 3427432997) % 7899123564644)

		fmt.Println(randNumber)

		for i := 0; i < stringLength; i++ {
			randNumber = uint64((211986194*randNumber + 5609872012) % 5609872012923)
			randLetter = uint64((894676913*randLetter + 3427432997) % 7899123564644)

			number := (randNumber % 9) + 48
			letter := (randLetter % 25) + 65

			if randNumber > randLetter {
				byteName.WriteByte(byte(number))
			} else if randNumber <= randLetter {
				byteName.WriteByte(byte(letter))
			}
		}

		g.Planet.Name = byteName.String()
	} else {
		g.Planet.Name = "Diamond Planet"
	}
}

func (g *Game) UpgradeStore(index int, levels int) string {
	if g.Gold.Cmp(g.Store[index].Cost) >= 0 {
		f := 1.03
		bulkCost := new(big.Float)

		for i := 1; i <= levels; i++ {
			pow := math.Pow(f, float64(g.Store[index].Level+int64(i-1)))
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

			g.CheckStoreLock(index)
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
			g.CheckStoreLock(k)
		}
	} else {
		return
	}

	g.CalculateShip(-1)
	g.CalculateCurrentDamage()
}

func (g *Game) CheckStoreLock(index int) {
	if index != -1 {
		if !g.Store[index].Locked {
			return
		}
		if g.Store[index].Level > 0 {
			if entry, ok := g.Store[index]; ok {
				entry.Locked = false
				g.Store[index] = entry
			}
		} else if g.Store[index].Level == 0 {
			cmp := g.Gold.Cmp(g.Store[index].BaseCost)
			if cmp >= 0 {
				if entry, ok := g.Store[index]; ok {
					entry.Locked = false
					g.Store[index] = entry
				}
			}
		}
	} else if index == -1 {
		for k := range g.Store {
			if !g.Store[k].Locked {
				continue
			}
			if g.Store[k].Level > 0 {
				if entry, ok := g.Store[k]; ok {
					entry.Locked = false
					g.Store[k] = entry
				}
			} else if g.Store[k].Level == 0 {
				cmp := g.Gold.Cmp(g.Store[k].BaseCost)
				if cmp >= 0 {
					if entry, ok := g.Store[k]; ok {
						entry.Locked = false
						g.Store[k] = entry
					}
				}
			}
		}
	}
}

func (g *Game) UpgradeShip(index int, levels int) string {
	if g.Gold.Cmp(g.Ship[index].Cost) >= 0 {
		if index == 3 {
			if g.Ship[index].Multiplier == 0.5 {
				return "max level reached"
			}
		}

		//f := 1.5
		bulkCost := new(big.Float)

		for i := 1; i <= levels; i++ {
			pow := math.Pow(g.Ship[index].Constant, float64(g.Ship[index].Level+int64(i-1)))
			bigPow := big.NewFloat(pow)
			cost := new(big.Float)

			cost.Mul(g.Ship[index].BaseCost, bigPow)

			bulkCost.Add(bulkCost, cost)
		}
		fmt.Println("   ---> bulkCost: ", bulkCost)

		if bulkCost.Cmp(g.Gold) > 0 {
			return "insufficient resources"
		}

		fmt.Printf("\n %v \n", bulkCost)

		if entry, ok := g.Ship[index]; ok {
			entry.Level += int64(levels)
			g.Ship[index] = entry
			result := new(big.Float)
			result.Sub(g.Gold, bulkCost) // g.Ship[index].Cost * levels
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

	if index == 1 {
		g.CalculateShipOne()
	} else if index == 2 {
		g.CalculateShipTwo()
	} else if index == 3 {
		g.CalculateShipThree()
	} else if index == 4 {
		g.CalculateShipFour()
	} else if index == -1 {
		g.CalculateShipOne()
		g.CalculateShipTwo()
		g.CalculateShipThree()
		g.CalculateShipFour()
	} else {
		return
	}

	g.CalculateCurrentDamage()
}

func (g *Game) UpgradeDiamondUpgrade(index int, levels int) string {
	if g.MaxLevel < 100 {
		return ""
	}

	if g.DiamondUpgrade[index].Cost.Cmp(g.Diamonds) > 0 {
		return "insufficient resources"
	}

	var bulkCost *big.Float
	for i := 1; i <= levels; i++ {
		pow := math.Pow(g.DiamondUpgrade[index].Constant, float64(g.DiamondUpgrade[index].Level+int64(i-1)))

		cost := big.NewFloat(float64(g.DiamondUpgrade[index].BaseCost) * pow)

		bulkCost.Add(bulkCost, cost)
	}

	if bulkCost.Cmp(g.Diamonds) > 0 {
		return "insufficient resources"
	}

	if entry, ok := g.DiamondUpgrade[index]; ok {
		entry.Level += int64(levels)
		g.DiamondUpgrade[index] = entry
		g.Diamonds.Sub(g.Diamonds, bulkCost)
	}
	g.Ch <- "upgrade"

	g.CalculateCurrentDamage()

	return ""
}

func (g *Game) CalculateDiamondUpgrade(index int) {
	if index != -1 {
		if entry, ok := g.DiamondUpgrade[index]; ok {
			pow := math.Pow(g.DiamondUpgrade[index].Constant, float64(g.DiamondUpgrade[index].Level))
			pow *= float64(g.DiamondUpgrade[index].BaseCost)

			result := new(big.Float).SetFloat64(pow)
			g.ConvertNumber(result, entry.Cost)

			if g.DiamondUpgrade[index].Level > 0 {
				bigPowMulInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(entry.Level), nil)
				bigPowMul := new(big.Float).SetInt(bigPowMulInt)
				bigPowMul.Mul(big.NewFloat(1.5), bigPowMul)
				entry.Multiplier = bigPowMul
			}
			g.DiamondUpgrade[index] = entry
		}
	} else if index == -1 {
		for k := range g.DiamondUpgrade {
			if entry, ok := g.DiamondUpgrade[k]; ok {
				pow := math.Pow(g.DiamondUpgrade[k].Constant, float64(g.DiamondUpgrade[k].Level))
				pow *= float64(g.DiamondUpgrade[k].BaseCost)

				result := new(big.Float).SetFloat64(pow)
				g.ConvertNumber(result, entry.Cost)

				if g.DiamondUpgrade[k].Level > 0 {
					bigPowMulInt := new(big.Int).Exp(big.NewInt(10), big.NewInt(entry.Level), nil)
					bigPowMul := new(big.Float).SetInt(bigPowMulInt)
					bigPowMul.Mul(big.NewFloat(1.5), bigPowMul)
					entry.Multiplier = bigPowMul
				}
				g.DiamondUpgrade[k] = entry
			}
		}
	}

	g.CalculateShip(index)
	g.CalculateCurrentDamage()
}

func (g *Game) CheckDiamondUpgradeUnlock() {
	if g.MaxLevel >= 100 {
		g.DiamondUpgradesUnlocked = true
	} else {
		g.DiamondUpgradesUnlocked = false
	}
}

func (g *Game) Advance() {
	if g.CurrentLevel == g.MaxLevel {
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
	}
	g.CheckBoss()
	g.CalculateDiamondPlanetChance()
	g.RollDiamondPlanet()
	g.GeneratePlanetName()
	fmt.Println("current stage: ", g.CurrentStage)
	fmt.Println("current level: ", g.CurrentLevel)
	fmt.Println("isboss?:", g.Planet.IsBoss)
}

func (g *Game) PreviousLevel() {
	if g.CurrentLevel > 1 {
		g.CurrentLevel -= 1
		g.CurrentStage = 10

		g.CheckBoss()

		g.RollDiamondPlanet()

		g.CalculatePlanetHealth()
		g.CalculateGoldEarned()

		g.GeneratePlanetName()
	}
}

func (g *Game) NextLevel() {
	if g.CurrentLevel < g.MaxLevel {
		g.CurrentLevel++
		if g.CurrentLevel == g.MaxLevel {
			g.CurrentStage = g.MaxStage
		} else {
			g.CurrentStage = 10
		}

		g.CheckBoss()

		g.RollDiamondPlanet()

		g.CalculatePlanetHealth()
		g.CalculateGoldEarned()

		g.GeneratePlanetName()
	}
}

func (g *Game) CalculateGoldEarned() {
	f := 1.35

	exp := float64(g.CurrentLevel)
	pow := math.Pow(f, exp)

	result := new(big.Float).SetFloat64(pow)
	if g.Planet.IsBoss && g.CurrentStage == 10 {
		result.Mul(result, big.NewFloat(10))
		if g.DiamondUpgrade[5].Level > 0 {
			result.Mul(result, g.DiamondUpgrade[5].Multiplier)
		}
	}

	bigMultiplier := big.NewFloat(g.Ship[4].Multiplier)
	if g.DiamondUpgrade[4].Level > 0 {
		bigMultiplier.Mul(bigMultiplier, g.DiamondUpgrade[4].Multiplier)
	}
	result.Mul(result, bigMultiplier)

	g.ConvertNumber(result, g.Planet.Gold)
}

func (g *Game) AddCurrentGold() {
	result := new(big.Float)
	result.Add(g.Gold, g.Planet.Gold)

	g.ConvertNumber(result, g.Gold)
}

func (g *Game) CalculateDiamondsEarned() {
	if g.Planet.IsBoss && g.MaxLevel > 99 && g.MaxLevel == g.CurrentLevel && !g.Planet.DiamondPlanet.IsDiamondPlanet {
		if g.MaxLevel == 100 {
			g.Planet.Diamonds = big.NewFloat(1) // .Add(g.Diamonds, big.NewFloat(1));
		} else {
			diamondConst := 1.1
			exp := float64(g.MaxLevel)
			pow := math.Pow(diamondConst, exp)

			result := new(big.Float).SetFloat64(pow)
			g.ConvertNumber(result, g.Planet.Diamonds)
			// g.Diamonds.Add(g.Diamonds, result) // int64((g.MaxLevel - 100) / 10)
		}
	} else if !g.Planet.IsBoss && g.MaxLevel > 99 && g.Planet.DiamondPlanet.IsDiamondPlanet {
		if g.MaxLevel == 100 {
			g.Planet.Diamonds = big.NewFloat(1)
		} else {
			diamondConst := 1.1
			exp := float64(g.MaxLevel)
			pow := math.Pow(diamondConst, exp)

			result := new(big.Float).SetFloat64(pow)
			result.Quo(result, big.NewFloat(2))
			g.ConvertNumber(result, g.Planet.Diamonds)
		}
	}
}

func (g *Game) AddCurrentDiamonds() {
	if g.Planet.IsBoss && g.MaxLevel > 99 && g.MaxLevel == g.CurrentLevel {
		result := new(big.Float)
		result.Add(g.Diamonds, g.Planet.Diamonds)

		g.ConvertNumber(result, g.Diamonds)
	}
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

// First index of ship map (1: DPS)
func (g *Game) CalculateShipOne() {
	// f := 1.5

	// cost = baseCost * 1.5 ** (level - 1)
	pow := math.Pow(g.Ship[1].Constant, float64(g.Ship[1].Level))
	bigPow := big.NewFloat(pow)

	cost := new(big.Float)
	cost.Mul(g.Ship[1].BaseCost, bigPow)
	g.ConvertNumber(cost, g.Ship[1].Cost)

	if entry, ok := g.Ship[1]; ok {
		// every 100ms
		entry.Multiplier = toFixed((0.1 * float64(g.Ship[1].Level)), 3) // 10
		g.Ship[1] = entry
	}

	//damage = currentDamage * multiplier
	bigMultiplier := big.NewFloat(g.Ship[1].Multiplier)
	g.Ship[1].Damage.Mul(g.CurrentDamage, bigMultiplier) // per 100ms
	if g.DiamondUpgrade[1].Level > 0 {
		g.Ship[1].Damage.Mul(g.Ship[1].Damage, g.DiamondUpgrade[1].Multiplier)
	}

	g.CheckShipLock(1)
}

// Second index of ship map (2: Click damage)
func (g *Game) CalculateShipTwo() {
	// f := 1.5

	// cost = baseCost * 1.5 ** (level - 1)
	pow := math.Pow(g.Ship[2].Constant, float64(g.Ship[2].Level))
	bigPow := big.NewFloat(pow)

	cost := new(big.Float)
	cost.Mul(g.Ship[2].BaseCost, bigPow)
	g.ConvertNumber(cost, g.Ship[2].Cost)

	if entry, ok := g.Ship[2]; ok {
		entry.Multiplier = toFixed((1.0 + 0.1*float64(g.Ship[2].Level)), 3)
		if g.DiamondUpgrade[2].Level > 0 {
			entry.Damage.Mul(g.CurrentDamage, g.DiamondUpgrade[2].Multiplier)
		} else {
			entry.Damage = g.CurrentDamage
		}
		g.Ship[2] = entry
	}

	//damage = currentDamage * multiplier
	//bigMultiplier := big.NewFloat(g.Ship[2].Multiplier)
	//g.Ship[2].Damage.Mul(g.CurrentDamage, bigMultiplier)
	// Since g.CurrentDamage is calculated in the other function
	// the g.Ship[2].Damage is going to be = g.CurrentDamage
	// in the if statement written above ^

	g.CheckShipLock(2)

}

// Third index of ship map (3: Critical Click)
func (g *Game) CalculateShipThree() {
	// f := 1.5

	// cost = baseCost * 1.5 ** (level - 1)
	pow := math.Pow(g.Ship[3].Constant, float64(g.Ship[3].Level))
	bigPow := big.NewFloat(pow)

	cost := new(big.Float)
	cost.Mul(g.Ship[3].BaseCost, bigPow)
	g.ConvertNumber(cost, g.Ship[3].Cost)

	if entry, ok := g.Ship[3]; ok {
		entry.Multiplier = toFixed((0.002 * float64(g.Ship[3].Level)), 3)
		if entry.Multiplier > 0.5 {
			entry.Multiplier = 0.5
		} // critical chance cap - 50%
		g.Ship[3] = entry
	}

	//critical damage = currentDamage * 5 <-- for now
	var bigMultiplier *big.Float
	if g.DiamondUpgrade[3].Level > 0 {
		bigMultiplier = new(big.Float)
		bigMultiplier.Mul(big.NewFloat(5.0), g.DiamondUpgrade[3].Multiplier)
	} else {
		bigMultiplier = big.NewFloat(5.0)
	}
	g.Ship[3].Damage.Mul(g.CurrentDamage, bigMultiplier)

	g.CheckShipLock(3)

}

// Fourth index of ship map (4: Planet gold)
func (g *Game) CalculateShipFour() {
	// f := 1.5

	// cost = baseCost * 1.5 ** (level - 1)
	pow := math.Pow(g.Ship[4].Constant, float64(g.Ship[4].Level))
	bigPow := big.NewFloat(pow)

	cost := new(big.Float)
	cost.Mul(g.Ship[4].BaseCost, bigPow)
	g.ConvertNumber(cost, g.Ship[4].Cost)

	if entry, ok := g.Ship[4]; ok {
		entry.Multiplier = toFixed((1.0 + 0.1*float64(g.Ship[4].Level)), 3)
		g.Ship[4] = entry
	}

	g.CalculateGoldEarned()

	g.CheckShipLock(4)
}

func (g *Game) CheckShipLock(index int) {
	if index != -1 {
		if !g.Ship[index].Locked {
			return
		}
		if g.Ship[index].Level == 0 {
			cmp := g.Gold.Cmp(g.Ship[index].BaseCost)
			if cmp >= 0 {
				if entry, ok := g.Ship[index]; ok {
					entry.Locked = false
					g.Ship[index] = entry
				}
			}
		} else if g.Ship[index].Level > 0 {
			if entry, ok := g.Ship[index]; ok {
				entry.Locked = false
				g.Ship[index] = entry
			}
		}
	} else if index == -1 {
		for k := range g.Ship {
			if !g.Ship[k].Locked {
				continue
			}
			if g.Ship[k].Level > 0 {
				if entry, ok := g.Ship[k]; ok {
					entry.Locked = false
					g.Ship[k] = entry
				}
			}

			if g.Ship[k].Level == 0 {
				cmp := g.Gold.Cmp(g.Ship[k].BaseCost)
				if cmp >= 0 {
					if entry, ok := g.Ship[k]; ok {
						entry.Locked = false
						g.Ship[k] = entry
					}
				}
			}
		}
	}
}

func (g *Game) CheckBoss() {
	if g.CurrentStage == 10 && g.CurrentLevel%10 == 0 {
		g.Planet.IsBoss = true
	} else {
		g.Planet.IsBoss = false
	}
}

func (g *Game) RollDiamondPlanet() {
	if !g.DiamondUpgradesUnlocked {
		return
	}
	if g.CurrentLevel%10 != 0 && !g.Planet.IsBoss {
		random := rand.Float64()
		if random <= g.Planet.DiamondPlanet.Chance {
			g.Planet.DiamondPlanet.IsDiamondPlanet = true
			g.Planet.Name = "Diamond Planet"
		} else {
			g.Planet.DiamondPlanet.IsDiamondPlanet = false
		}
	}

	fmt.Println(g.Planet.Name)
	fmt.Println("IsDiamondPlanet", g.Planet.DiamondPlanet.IsDiamondPlanet)
}

func (g *Game) CalculateDiamondPlanetChance() {
	if !g.DiamondUpgradesUnlocked {
		return
	}
	g.Planet.DiamondPlanet.Chance = g.Planet.DiamondPlanet.BaseChance + (0.00001 * (float64(g.MaxLevel-100) / 10))
}

func toFixed(num float64, prec int) float64 {
	output := math.Pow(10, float64(prec))
	round := int((num * output) + math.Copysign(0.5, num))
	return float64(round) / output
}
