package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
)

var Bind *sqlx.DB

func init() {
	var err error
	Bind, err = sqlx.Open("sqlite3", "./resources/db.sqlite3?_busy_timeout=5000")
	if err != nil {
		panic(err)
	}

	_, err = Bind.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		panic(err)
	}
}

func Migrate() error {
	m, err := migrate.New(
		"file://./db/migrations",
		fmt.Sprintf("sqlite3://%s", "./resources/db.sqlite3"))
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
