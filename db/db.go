package db

import (
	"gopkg.in/pg.v4"
)

var database *pg.DB

// PostgresCfg cfg for postgresql
type PostgresCfg struct {
	Address  string `toml:"address"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

// InitDB new database then connect
// Just Simple use password
func InitDB(cfg PostgresCfg) {
	database = pg.Connect(&pg.Options{
		Addr:     cfg.Address,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})
}

// Get return database interface as it is corountine safe
func Get() *pg.DB {
	return database
}

// Close close db connect
func Close() error {
	return database.Close()
}
