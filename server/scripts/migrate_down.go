package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDown() {
	m, err := migrate.New(
		"file://./db/migrations",
		fmt.Sprintf("sqlite3://%s", "./db.sqlite3"))
	if err != nil {
		fmt.Printf("create migrate failed: %v\n", err)
		return
	}

	if err := m.Down(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("migrate no change")
			return
		}
		fmt.Printf("migrate down failed: %v\n", err)
		return
	}

	fmt.Println("migrate down success")
}
