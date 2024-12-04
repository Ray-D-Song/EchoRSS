#!/bin/bash

DB_PATH="./server/resources/db.sqlite3"
MIGRATION_DIR="./server/db/migrations"

# check database file exists
if [ ! -f "$DB_PATH" ]; then
    echo "Database file not found"
    exit 1
fi

# get latest up migration file
latest_migration=$(ls -1 "$MIGRATION_DIR"/*up.sql 2>/dev/null | sort -r | head -n 1)

if [ -z "$latest_migration" ]; then
    echo "migrate no change"
    exit 0
fi

# execute migration
if sqlite3 "$DB_PATH" < "$latest_migration"; then
    echo "migrate up success"
else
    echo "migrate up failed"
    exit 1
fi