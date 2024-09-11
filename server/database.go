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

func createShips(p *Postgres) error {
	createShips := `CREATE TABLE IF NOT EXISTS ships (
		id INT,
		level INT,
		multiplier FLOAT(4),
		PRIMARY KEY (id)
	)`

	_, err := p.db.Exec(createShips)
	return err
}

func createStores(p *Postgres) error {
	createStores := `CREATE TABLE IF NOT EXISTS stores (
		id INT,
		level INT,
		damage VARCHAR(50),
		PRIMARY KEY (id)	
	)`

	_, err := p.db.Exec(createStores)
	return err
}

func createGames(p *Postgres) error {
	createGames := `CREATE TABLE IF NOT EXISTS games (
		id INT,
		gold VARCHAR(50),
		diamonds VARCHAR(50),
		current_damage VARCHAR(50),
		max_damage VARCHAR(50),
		current_level INT,
		max_level INT,
		current_stage SMALLINT,
		max_stage SMALLINT,
		planets_destroyed INT,
		store_id INT,
		ship_id INT,
		PRIMARY KEY (id),
		FOREIGN KEY(store_id)
		REFERENCES stores(id),
		FOREIGN KEY(ship_id)
		REFERENCES ships(id)
	)`

	_, err := p.db.Exec(createGames)
	return err
}

func createUsers(p *Postgres) error {
	createUsers := `CREATE TABLE IF NOT EXISTS users (
		id INT,
		login VARCHAR(50),
		email VARCHAR(50),
		password VARCHAR(100),
		game_id INT,
		PRIMARY KEY (id),
		FOREIGN KEY(game_id)
		REFERENCES games(id)
	)`

	_, err := p.db.Exec(createUsers)
	return err
}

func (p *Postgres) PrepareDb() error {
	err := createShips(p)
	if err != nil {
		return err
	}

	err = createStores(p)
	if err != nil {
		return err
	}

	err = createGames(p)
	if err != nil {
		return err
	}

	err = createUsers(p)
	if err != nil {
		return err
	}

	return err
}
