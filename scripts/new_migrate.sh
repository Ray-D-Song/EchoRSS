#!/bin/bash

read -p "migrate description: " description

# generate timestamp
timestamp=$(date '+%Y%m%d%H%M%S')

# set migration dir
migration_dir="./server/db/migrations"

# create migration dir if not exists
mkdir -p "$migration_dir"

# generate migration file names
up_file="${migration_dir}/${timestamp}_${description}.up.sql"
down_file="${migration_dir}/${timestamp}_${description}.down.sql"

# create migration files
echo "-- SQL" > "$up_file"
echo "-- SQL" > "$down_file"

# output created file names
echo "create migration files:"
echo "$timestamp"_"$description.up.sql"
echo "$timestamp"_"$description.down.sql"
