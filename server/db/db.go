package db

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
)

var Bind *sqlx.DB
var dbPath string

func init() {
	// check if './db.sqlite3' exists
	if _, err := os.Stat("./db.sqlite3"); os.IsNotExist(err) {
		// create db.sqlite3
		_, err := os.Create("./db.sqlite3")
		if err != nil {
			panic(err)
		}
	}
	dbPath = "./db.sqlite3"
	var err error
	Bind, err = sqlx.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
}

func Migrate() error {
	m, err := migrate.New(
		"file://./db/migrations",
		fmt.Sprintf("sqlite3://%s", dbPath))
	if err != nil {
		fmt.Printf("create migrate failed: %v\n", err)
		return err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("migrate no change")
			return nil
		}
		return err
	}

	return nil
}
