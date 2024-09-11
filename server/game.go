package main

type ShipUpgrade struct {
	Level      int
	Multiplier float64
	Damage     int //!later to be edited in the database
}

type Ship struct {
	Upgrades map[string]ShipUpgrade
}

type StoreUpgrade struct {
	Level  int
	Damage int
}

type Store struct {
	Upgrades map[string]StoreUpgrade
}
