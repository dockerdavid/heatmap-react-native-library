package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Client struct {
	DB *sql.DB
}

func NewClient(MobileDir string) (*Client, error) {
	err := os.MkdirAll(MobileDir+"/db", os.ModePerm)

	if err != nil {
		return &Client{}, err
	}

	_, err = os.Create(MobileDir + "/db/client.db")

	if err != nil {
		return &Client{}, err
	}

	database, err := sql.Open("sqlite3", MobileDir+"/db/client.db")

	if err != nil {
		return &Client{}, err
	}

	prepare, err := database.Prepare(`CREATE TABLE IF NOT EXISTS capture (
			coordinate_x REAL,
			coordinate_y REAL,
			screen_name TEXT,
			uuid TEXT,
			image TEXT,
			status TEXT,
			created_at TEXT,
			updated_at TEXT
		);`)

	if err != nil {
		return &Client{}, err
	}

	_, err = prepare.Exec()

	if err != nil {
		return &Client{}, err
	}

	prepare, err = database.Prepare("CREATE TABLE IF NOT EXISTS client (api_key TEXT PRIMARY KEY)")

	if err != nil {
		return &Client{}, err
	}

	_, err = prepare.Exec()

	if err != nil {
		return &Client{}, err
	}

	prepare, err = database.Prepare(`CREATE TABLE IF NOT EXISTS device (
			id TEXT PRIMARY KEY,
			name TEXT,
			brand TEXT,
			model TEXT,
			manufacturer TEXT,
			screen_width REAL,
			screen_height REAL,
			build_id TEXT,
			battery_level REAL,
			build_number TEXT,
			bundle_id TEXT,
			free_disk_storage REAL,
			power_state TEXT,
			readable_version TEXT,
			system_name TEXT,
			system_version TEXT
		);`)

	if err != nil {
		return &Client{}, err
	}

	_, err = prepare.Exec()

	if err != nil {
		return &Client{}, err
	}

	return &Client{DB: database}, nil
}
