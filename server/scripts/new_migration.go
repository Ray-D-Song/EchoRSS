package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func NewMigration() {
	fmt.Print("migrate description: ")
	var description string
	fmt.Scanln(&description)

	timestamp := time.Now().Format("20060102150405")

	upFileName := fmt.Sprintf("%s_%s.up.sql", timestamp, description)
	downFileName := fmt.Sprintf("%s_%s.down.sql", timestamp, description)

	migrationDir := "./db/migrations"

	if err := os.MkdirAll(migrationDir, 0755); err != nil {
		fmt.Printf("create dir failed: %v\n", err)
		return
	}

	upFile := filepath.Join(migrationDir, upFileName)
	if err := os.WriteFile(upFile, []byte("-- SQL\n"), 0644); err != nil {
		fmt.Printf("create up file failed: %v\n", err)
		return
	}

	downFile := filepath.Join(migrationDir, downFileName)
	if err := os.WriteFile(downFile, []byte("-- SQL\n"), 0644); err != nil {
		fmt.Printf("create down file failed: %v\n", err)
		return
	}

	fmt.Printf("create migration files:\n%s\n%s\n", upFileName, downFileName)
}
