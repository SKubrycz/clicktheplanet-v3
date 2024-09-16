package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

type Database interface {
	CreateAccount(*User) error
	GetAccountByLogin(string) (*User, error)
	GetGameByUserId(id int) (*GameData, error)
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

	return &Postgres{
		db: db,
	}, nil
}

// Function for /register route
// Inserts into every table for game initialization
func (p *Postgres) CreateAccount(u *User) error {
	queryUser := `INSERT INTO users (login, email, password, created_at) VALUES ($1, $2, $3, $4)`
	queryGame := `
	INSERT INTO games (gold, diamonds, max_damage, current_level, max_level, current_stage, max_stage, planets_destroyed, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;
	`
	queryGameShip := `
	INSERT INTO game_ship (level, game_id, ship_id) VALUES ($1, $2, $3);
	`

	queryGameStore := `
	INSERT INTO game_store (level, game_id, store_id) VALUES ($1, $2, $3);
	`

	_, err := p.db.Exec(queryUser, u.Login, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}

	user, err := p.GetAccountByLogin(u.Login)
	if err != nil {
		return err
	}

	var gameId int
	err = p.db.QueryRow(queryGame, 100, 0, 1, 1, 1, 1, 1, 0, user.Id).Scan(&gameId)
	if err != nil {
		return err
	}

	for i := 1; i <= 4; i++ {
		_, err = p.db.Exec(queryGameShip, 0, gameId, i)
		if err != nil {
			fmt.Println(err)
			return err
		}

		_, err = p.db.Exec(queryGameStore, 0, gameId, i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (p *Postgres) GetAccountByLogin(login string) (*User, error) {
	query := `SELECT * FROM users WHERE login = $1`

	rows, err := p.db.Query(query, login)
	if err != nil {
		return nil, err
	}

	var user *User
	for rows.Next() {
		user = new(User)
		err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)
	}
	fmt.Println(user)
	return user, err
}

func (p *Postgres) GetGameByUserId(id int) (*GameData, error) {
	gameQuery := `
	SELECT id, gold, diamonds, max_damage, current_level, max_level, current_stage, max_stage, planets_destroyed
	FROM games WHERE user_id = $1
	`

	gameShipQuery := `
	SELECT level FROM game_ship WHERE game_id = $1
	`

	gameStoreQuery := `
	SELECT level FROM game_store WHERE game_id = $1
	`

	var game = &GameData{
		Store: map[string]StoreUpgradeData{},
		Ship:  map[string]ShipUpgradeData{},
	}
	var gameId int
	err := p.db.QueryRow(gameQuery, id).Scan(
		&gameId,
		&game.Gold,
		&game.Diamonds,
		&game.MaxDamage,
		&game.CurrentLevel,
		&game.MaxLevel,
		&game.CurrentStage,
		&game.MaxStage,
		&game.PlanetsDestroyed,
	)
	if err != nil {
		return nil, err
	}

	i := 1
	rows, err := p.db.Query(gameShipQuery, gameId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		ship := new(ShipUpgradeData)
		err := rows.Scan(
			&ship.Level,
		)
		if err != nil {
			return nil, err
		}
		iStr := strconv.Itoa(i)
		game.Ship[iStr] = *ship
		i++
	}

	i = 1
	rows, err = p.db.Query(gameStoreQuery, gameId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		store := new(StoreUpgradeData)
		err := rows.Scan(
			&store.Level,
		)
		if err != nil {
			return nil, err
		}
		iStr := strconv.Itoa(i)
		fmt.Println(iStr)
		game.Store[iStr] = *store
		i++
	}

	return game, nil
}
