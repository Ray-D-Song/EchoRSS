package utils

import (
	"os"
	"path"
)

func EnsureDir() error {
	// ensure ./resources
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	if _, err := os.Stat(path.Join(dir, "resources")); os.IsNotExist(err) {
		err = os.Mkdir(path.Join(dir, "resources"), 0755)
		if err != nil {
			return err
		}
	}
	// ensure ./resources/logs
	if _, err := os.Stat(path.Join(dir, "resources", "logs")); os.IsNotExist(err) {
		err = os.Mkdir(path.Join(dir, "resources", "logs"), 0755)
		if err != nil {
			return err
		}
	}
	// ensure ./resources/logs/server.log
	if _, err := os.Stat(path.Join(dir, "resources", "logs", "server.log")); os.IsNotExist(err) {
		_, err = os.Create(path.Join(dir, "resources", "logs", "server.log"))
		if err != nil {
			return err
		}
	}

	// ensure ./resources/db.sqlite3
	if _, err := os.Stat(path.Join(dir, "resources", "db.sqlite3")); os.IsNotExist(err) {
		_, err = os.Create(path.Join(dir, "resources", "db.sqlite3"))
		if err != nil {
			return err
		}
	}

	return nil
}
