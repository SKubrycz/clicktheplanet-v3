package main

type RegisterCredentials struct {
	Login         string `json:"login"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AgainPassword string `json:"againPassword"`
}

type LoginCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserId string

type UserData struct {
	Login string `json:"login"`
	// also later image
}

type ShipUpgradeData struct {
	Level      int64
	Multiplier float64
	Damage     string
}

type StoreUpgradeData struct {
	Level  int64
	Damage string
}

type DiamondUpgradeUpgradeData struct {
	Level int64
}

// For data retrieved from the database
type GameData struct {
	Id               int64
	Gold             string
	Diamonds         string
	MaxDamage        string
	CurrentLevel     int64
	MaxLevel         int64
	CurrentStage     uint8
	MaxStage         uint8
	PlanetsDestroyed string
	Store            map[int]StoreUpgradeData
	Ship             map[int]ShipUpgradeData
	DiamondUpgrade   map[int]DiamondUpgradeUpgradeData
}
