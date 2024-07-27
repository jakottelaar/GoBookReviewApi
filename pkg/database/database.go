package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jakottelaar/gobookreviewapp/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Initialize(cfg *config.Config) error {

	var err error

	DB, err = sql.Open("postgres", cfg.Database.Dsn)

	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}

	DB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	DB.SetConnMaxIdleTime(cfg.Database.MaxIdleTime)

	log.Println("Successfully connected to the database")

	return nil

}

func GetDB() *sql.DB {
	return DB
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
