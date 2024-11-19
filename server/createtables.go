package main

func createShipTable(p *Postgres) error {
	createShip := `CREATE TABLE IF NOT EXISTS ship (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50),
		description TEXT
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_title_ship ON ship (title);
	`

	_, err := p.db.Exec(createShip)
	return err
}

func createStoreTable(p *Postgres) error {
	createStore := `CREATE TABLE IF NOT EXISTS store (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50),
		description TEXT
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_title_store ON store (title);
	`

	_, err := p.db.Exec(createStore)
	return err
}

func createDiamondUpgradeTable(p *Postgres) error {
	createDiamondUpgrade := `CREATE TABLE IF NOT EXISTS diamond_upgrade (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50),
		description TEXT
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_title_diamond_upgrade ON diamond_upgrade (title);
	`

	_, err := p.db.Exec(createDiamondUpgrade)
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

// TODO: Recalculate current_damage during runtime after database fetch
func createGamesTable(p *Postgres) error {
	createGames := `CREATE TABLE IF NOT EXISTS games (
		id SERIAL PRIMARY KEY,
		gold VARCHAR(50),
		diamonds VARCHAR(50),
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

// "Dictionary" table for a singular user account
// and whether he bought a upgrade in the Ship Tab
// essentially, a table containing information of each
// upgrade and it's availability for the player (Lock/Unlocked)
// --- the multiplier will be calculated using a formula dependent on a level during runtime
func createGameShipTable(p *Postgres) error {
	createGameShip := `CREATE TABLE IF NOT EXISTS game_ship (
		id SERIAL PRIMARY KEY,
		level BIGINT,
		diamond_level INT,
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

func createGameStoreTable(p *Postgres) error {
	createGameStore := `CREATE TABLE IF NOT EXISTS game_store(
		id SERIAL PRIMARY KEY,
		level BIGINT,
		game_id INT NOT NULL,
		store_id INT NOT NULL,
		FOREIGN KEY (game_id)
		REFERENCES games(id),
		FOREIGN KEY (store_id)
		REFERENCES store(id)
	)`

	_, err := p.db.Exec(createGameStore)
	return err
}

func createGameDiamondUpgradeTable(p *Postgres) error {
	createGameDiamondUpgrade := `CREATE TABLE IF NOT EXISTS game_diamond_upgrade (
		id SERIAL PRIMARY KEY,
		level BIGINT,
		game_id INT NOT NULL,
		diamond_upgrade_id INT NOT NULL,
		FOREIGN KEY (game_id)
		REFERENCES games(id),
		FOREIGN KEY (diamond_upgrade_id)
		REFERENCES diamond_upgrade(id)
	)`

	_, err := p.db.Exec(createGameDiamondUpgrade)
	return err
}

func insertShipTable(p *Postgres) error {
	insertShip := `
	INSERT INTO ship (id, title, description) VALUES (1, 'Upgrade 1', 'Upgrade 1 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO ship (id, title, description) VALUES (2, 'Upgrade 2', 'Upgrade 2 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO ship (id, title, description) VALUES (3, 'Upgrade 3', 'Upgrade 3 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO ship (id, title, description) VALUES (4, 'Upgrade 4', 'Upgrade 4 desc') ON CONFLICT (title) DO NOTHING;
	`

	_, err := p.db.Exec(insertShip)
	return err
}

func insertStoreTable(p *Postgres) error {
	insertStore := `
	INSERT INTO store (id, title, description) VALUES (1, 'Store 1', 'Store 1 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (2, 'Store 2', 'Store 2 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (3, 'Store 3', 'Store 3 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (4, 'Store 4', 'Store 4 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (5, 'Store 5', 'Store 5 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (6, 'Store 6', 'Store 6 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (7, 'Store 7', 'Store 7 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (8, 'Store 8', 'Store 8 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (9, 'Store 9', 'Store 9 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (10, 'Store 10', 'Store 10 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (11, 'Store 11', 'Store 11 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (12, 'Store 12', 'Store 12 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (13, 'Store 13', 'Store 13 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (14, 'Store 14', 'Store 14 desc') ON CONFLICT (title) DO NOTHING;
	INSERT INTO store (id, title, description) VALUES (15, 'Store 15', 'Store 15 desc') ON CONFLICT (title) DO NOTHING;
	`

	_, err := p.db.Exec(insertStore)
	return err
}

func insertDiamondUpgradeTable(p *Postgres) error {
	insertDiamondUpgrade := `
	INSERT INTO diamond_upgrade (id, title, description) VALUES (1, 'Dps', 'Enhance idle damage') ON CONFLICT (title) DO NOTHING;
	INSERT INTO diamond_upgrade (id, title, description) VALUES (2, 'Click damage', 'Enhance current click damage') ON CONFLICT (title) DO NOTHING;
	INSERT INTO diamond_upgrade (id, title, description) VALUES (3, 'Critical damage', 'Multiply critical click damage') ON CONFLICT (title) DO NOTHING;
	INSERT INTO diamond_upgrade (id, title, description) VALUES (4, 'Planet gold', 'Enhance gold gained from Planets') ON CONFLICT (title) DO NOTHING;
	INSERT INTO diamond_upgrade (id, title, description) VALUES (5, 'Boss gold', 'Enhance gold gained from Boss levels') ON CONFLICT (title) DO NOTHING;
	`

	_, err := p.db.Exec(insertDiamondUpgrade)
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

	err = createDiamondUpgradeTable(p)
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

	err = createGameStoreTable(p)
	if err != nil {
		return err
	}

	err = createGameDiamondUpgradeTable(p)
	if err != nil {
		return err
	}

	//INSERTS
	err = insertShipTable(p)
	if err != nil {
		return err
	}

	err = insertStoreTable(p)
	if err != nil {
		return err
	}

	err = insertDiamondUpgradeTable(p)
	if err != nil {
		return err
	}

	return err
}
