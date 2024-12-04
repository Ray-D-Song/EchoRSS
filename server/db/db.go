package db

import (
	"embed"
	"fmt"
	"os"
	"path"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	"ray-d-song.com/echo-rss/utils"
)

//go:embed migrations/*.sql
var schemaFs embed.FS

var Bind *sqlx.DB

func init() {
	utils.EnsureDir()
	var err error
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dsn := path.Join(dir, "resources", "db.sqlite3?_busy_timeout=5000&_journal_mode=WAL&_sync=NORMAL")
	Bind, err = sqlx.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
}

func Migrate() error {
	d, err := iofs.New(schemaFs, "migrations")
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithSourceInstance(
		"iofs", d,
		fmt.Sprintf("sqlite3://%s", "./resources/db.sqlite3"))
	if err != nil {
		panic(err)
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
