package main

import "os"

func main() {
	exec := os.Args[1]
	switch exec {
	case "new":
		NewMigration()
	case "down":
		MigrateDown()
	}
}
