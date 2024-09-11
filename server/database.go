package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database interface {
}

type Postgres struct {
	db *sql.DB
}

func NewDatabase() (*Postgres, error) {
	conn := "user=postgres dbname=clicktheplanet-v3 password=admin sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")

	return &Postgres{
		db: db,
	}, nil
}

func createShipTable(p *Postgres) error {
	createShip := `CREATE TABLE IF NOT EXISTS ship (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50),
		description TEXT
	);`

	_, err := p.db.Exec(createShip)
	return err
}

// "Dictionary" table for a singular user account
// and whether he bought a upgrade in the Ship Tab
// essentially, a table containing information of each
// / upgrade and it's availability for the player (Lock/Unlocked)
func createGameShipTable(p *Postgres) error {
	createGameShip := `CREATE TABLE IF NOT EXISTS game_ship (
		id SERIAL PRIMARY KEY,
		level INT,
		multiplier FLOAT(4),
		damage VARCHAR(50),
		game_id INT NOT NULL,
		ship_id INT NOT NULL,
		FOREIGN KEY (game_id)
		REFERENCES games(id),
		FOREIGN KEY (ship_id)
		REFERENCES ship(id)
	)`

	_, err := p.db.Exec(createGameShip)
	return err

}

func createStoreTable(p *Postgres) error {
	createStore := `CREATE TABLE IF NOT EXISTS store (
		id SERIAL PRIMARY KEY,
		level INT,
		damage VARCHAR(50)
	)`

	_, err := p.db.Exec(createStore)
	return err
}

func createUsersTable(p *Postgres) error {
	createUsers := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(50),
		email VARCHAR(50),
		password VARCHAR(100),
		created_at TIMESTAMP
	)`

	_, err := p.db.Exec(createUsers)
	return err
}

func createGamesTable(p *Postgres) error {
	createGames := `CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		gold VARCHAR(50),
		diamonds VARCHAR(50),
		current_damage VARCHAR(50),
		max_damage VARCHAR(50),
		current_level INT,
		max_level INT,
		current_stage SMALLINT,
		max_stage SMALLINT,
		planets_destroyed INT,
		user_id INT NOT NULL,
		FOREIGN KEY(user_id)
		REFERENCES users(id)
	)`

	_, err := p.db.Exec(createGames)
	return err
}

func (p *Postgres) PrepareDb() error {
	err := createShipTable(p)
	if err != nil {
		return err
	}

	err = createStoreTable(p)
	if err != nil {
		return err
	}

	err = createUsersTable(p)
	if err != nil {
		return err
	}

	err = createGamesTable(p)
	if err != nil {
		return err
	}

	err = createGameShipTable(p)
	if err != nil {
		return err
	}

	//later add some INSERTs to database to fill with initial game data available for all users

	return err
}

// Function for /register route
// Creates every table connected with foreign keys for a new account
func (p *Postgres) createAccount(u *User) error {
	return nil
}
