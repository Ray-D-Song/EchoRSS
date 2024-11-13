package utils

import "os"

func EnsureDir() error {
	// ensure ./resources
	if _, err := os.Stat("./resources"); os.IsNotExist(err) {
		err = os.Mkdir("./resources", 0755)
		if err != nil {
			return err
		}
	}
	// ensure ./resources/logs
	if _, err := os.Stat("./resources/logs"); os.IsNotExist(err) {
		err = os.Mkdir("./resources/logs", 0755)
		if err != nil {
			return err
		}
	}
	// ensure ./resources/logs/server.log
	if _, err := os.Stat("./resources/logs/server.log"); os.IsNotExist(err) {
		_, err = os.Create("./resources/logs/server.log")
		if err != nil {
			return err
		}
	}

	// ensure ./resources/db.sqlite3
	if _, err := os.Stat("./resources/db.sqlite3"); os.IsNotExist(err) {
		_, err = os.Create("./resources/db.sqlite3")
		if err != nil {
			return err
		}
	}

	return nil
}
