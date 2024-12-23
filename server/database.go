package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database interface {
	CreateAccount(*User) error
	DeleteAccountById(id int) error
	GetAccountByLogin(string) (*User, error)
	GetAccountById(id int) (*User, error)
	GetGameByUserId(id int) (*GameData, error)
	SaveGameProgress(userId int, g *Game) error
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
	fmt.Println("Creating account...")

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

	queryGameDiamondUpgrade := `
	INSERT INTO game_diamond_upgrade (level, game_id, diamond_upgrade_id) VALUES ($1, $2, $3);
	`

	queryCountShip := `
	SELECT COUNT(id) FROM ship;
	`

	queryCountStore := `
	SELECT count(id) FROM store;
	`

	queryCountDiamondUpgrade := `
	SELECT count(id) from diamond_upgrade;
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
	err = p.db.QueryRow(queryGame, 0, 0, 1, 1, 1, 1, 1, 0, user.Id).Scan(&gameId)
	if err != nil {
		return err
	}

	//Count rows for store and ship before inserting to game_ship and game_store
	var shipCount int
	err = p.db.QueryRow(queryCountShip).Scan(&shipCount)
	if err != nil {
		return err
	}
	var storeCount int
	err = p.db.QueryRow(queryCountStore).Scan(&storeCount)
	if err != nil {
		return err
	}
	var diamondUpgradeCount int
	err = p.db.QueryRow(queryCountDiamondUpgrade).Scan(&diamondUpgradeCount)
	if err != nil {
		return err
	}

	for i := 1; i <= shipCount; i++ {
		_, err = p.db.Exec(queryGameShip, 0, gameId, i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	for i := 1; i <= storeCount; i++ {
		_, err = p.db.Exec(queryGameStore, 0, gameId, i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	for i := 1; i <= diamondUpgradeCount; i++ {
		_, err = p.db.Exec(queryGameDiamondUpgrade, 0, gameId, i)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (p *Postgres) DeleteAccountById(id int) error {
	querySelect := `SELECT id FROM games WHERE user_id = $1`

	deleteFromGameStore := `DELETE FROM game_store WHERE game_id = $1`
	deleteFromGameShip := `DELETE FROM game_ship WHERE game_id = $1`
	deleteFromGameDiamondUpgrade := `DELETE FROM game_diamond_upgrade WHERE game_id = $1`
	deleteFromGames := `DELETE FROM games WHERE user_id = $1`
	deleteFromUsers := `DELETE FROM users WHERE id = $1`

	var gameId int

	err := p.db.QueryRow(querySelect, id).Scan(&gameId)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(deleteFromGameStore, gameId)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(deleteFromGameShip, gameId)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(deleteFromGameDiamondUpgrade, gameId)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(deleteFromGames, id)
	if err != nil {
		return err
	}

	_, err = p.db.Exec(deleteFromUsers, id)
	if err != nil {
		return err
	}

	fmt.Printf("\nUser %v deleted\n", id)

	return nil
}

func (p *Postgres) GetAccountByLogin(login string) (*User, error) {
	query := `SELECT * FROM users WHERE login = $1`

	fmt.Println("Searching for existing user...")
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
	fmt.Println("User result: ", user)
	return user, err
}

func (p *Postgres) GetAccountById(id int) (*User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	fmt.Println("Searching for existing user...")
	rows, err := p.db.Query(query, id)
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
	fmt.Println("User result: ", user)
	return user, err
}

func (p *Postgres) GetGameByUserId(id int) (*GameData, error) {
	gameQuery := `
	SELECT id, gold, diamonds, max_damage, current_level, max_level, current_stage, max_stage, planets_destroyed
	FROM games WHERE user_id = $1
	`

	storeCountQuery := `
	SELECT COUNT(id) FROM store;
	`

	gameStoreCountQuery := `
	SELECT count(store_id) FROM game_store WHERE game_id = $1 GROUP BY game_id;
	`

	gameStoreInsertQuery := `
	INSERT INTO game_store (level, game_id, store_id) VALUES ($1, $2, $3);
	`

	gameShipQuery := `
	SELECT level FROM game_ship WHERE game_id = $1 ORDER BY id ASC
	`

	gameStoreQuery := `
	SELECT level FROM game_store WHERE game_id = $1 ORDER BY id ASC
	`

	gameDiamondUpgradeQuery := `
	SELECT level FROM game_diamond_upgrade WHERE game_id = $1 ORDER BY id ASC
	`

	var game = &GameData{
		Store:          map[int]StoreUpgradeData{},
		Ship:           map[int]ShipUpgradeData{},
		DiamondUpgrade: map[int]DiamondUpgradeUpgradeData{},
	}
	err := p.db.QueryRow(gameQuery, id).Scan(
		&game.Id,
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

	var storeCount int
	err = p.db.QueryRow(storeCountQuery).Scan(&storeCount)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var gameStoreCount int
	err = p.db.QueryRow(gameStoreCountQuery, game.Id).Scan(&gameStoreCount)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	storeDiff := storeCount - gameStoreCount

	fmt.Println("STORE DIFF:", storeDiff)

	if storeDiff != 0 {
		for i := gameStoreCount + 1; i <= storeCount; i++ {
			fmt.Println(i)
			fmt.Println(gameStoreCount)
			_, err = p.db.Exec(gameStoreInsertQuery, 0, game.Id, i)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		}
	}

	i := 1
	rows, err := p.db.Query(gameShipQuery, game.Id)
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
		game.Ship[i] = *ship
		i++
	}

	i = 1
	rows, err = p.db.Query(gameStoreQuery, game.Id)
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
		game.Store[i] = *store
		i++
	}

	i = 1
	rows, err = p.db.Query(gameDiamondUpgradeQuery, game.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		diamondUpgrade := new(DiamondUpgradeUpgradeData)
		err := rows.Scan(
			&diamondUpgrade.Level,
		)
		if err != nil {
			return nil, err
		}
		game.DiamondUpgrade[i] = *diamondUpgrade
		i++
	}

	return game, nil
}

func (p *Postgres) SaveGameProgress(userId int, g *Game) error {
	queryGame := `
	UPDATE games
	SET gold = $1, diamonds = $2,
	max_damage = $3,
	current_level = $4, max_level = $5,
	current_stage = $6, max_stage = $7,
	planets_destroyed = $8
	WHERE user_id = $9
	`
	// Also, save store and ship upgrades
	queryGameShip := `
		UPDATE game_ship
		SET level = $1
		WHERE game_id = $2 AND ship_id = $3;
	`

	queryGameStore := `
		UPDATE game_store
		SET level = $1
		WHERE game_id = $2 AND store_id = $3;
	`

	queryGameDiamondUpgrade := `
		UPDATE game_diamond_upgrade
		SET level = $1
		WHERE game_id = $2 AND diamond_upgrade_id = $3;
	`

	_, err := p.db.Exec(queryGame,
		g.Gold.String(),
		g.Diamonds.String(),
		g.MaxDamage.String(),
		g.CurrentLevel,
		g.MaxLevel,
		g.CurrentStage,
		g.MaxStage,
		g.PlanetsDestroyed.String(),
		userId,
	)
	if err != nil {
		return err
	}

	for i := 1; i <= len(g.Ship); i++ {
		_, err := p.db.Exec(queryGameShip, g.Ship[i].Level, g.Id, i)
		if err != nil {
			return err
		}
	}

	for i := 1; i <= len(g.Store); i++ {
		_, err := p.db.Exec(queryGameStore, g.Store[i].Level, g.Id, i)
		if err != nil {
			return err
		}
	}

	for i := 1; i <= len(g.DiamondUpgrade); i++ {
		_, err := p.db.Exec(queryGameDiamondUpgrade, g.DiamondUpgrade[i].Level, g.Id, i)
		if err != nil {
			return err
		}
	}

	fmt.Println("Saved game")
	return nil
}
