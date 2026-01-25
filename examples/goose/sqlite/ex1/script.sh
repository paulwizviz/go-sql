#!/bin/bash

if [ "$(basename $(realpath .))" != "go-sql" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export GOOSE_MIGRATION_DIR=./examples/goose/sqlite/ex1/migrations

COMMAND="$1"

case $COMMAND in
    "down")
        goose sqlite3 ./sqlite.db down # Drop constrain
        sleep 2
        goose sqlite3 ./sqlite.db down # Drop columns
        sleep 2
        goose sqlite3 ./sqlite.db down # Drop tables
        ;;
    "status")
        goose sqlite3 ./sqlite.db status
        ;;
    "up")
        goose sqlite3 ./sqlite.db up
        ;;
    *)
        echo "Usage: $0 [down | status | up]"
        ;;
esac