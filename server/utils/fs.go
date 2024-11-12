package utils

import "os"

func EnsureDir() error {
	// ensure ./resources
	if _, err := os.Stat("./resources"); os.IsNotExist(err) {
		return os.Mkdir("./resources", 0755)
	}
	// ensure ./resources/logs
	if _, err := os.Stat("./resources/logs"); os.IsNotExist(err) {
		return os.Mkdir("./resources/logs", 0755)
	}
	// ensure ./resources/logs/server.log
	if _, err := os.Stat("./resources/logs/server.log"); os.IsNotExist(err) {
		_, err = os.Create("./resources/logs/server.log")
		if err != nil {
			return err
		}
	}
	return nil
}
