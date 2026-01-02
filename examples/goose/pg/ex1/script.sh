#!/bin/bash

if [ "$(basename $(realpath .))" != "go-sql" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgres://postgres:postgres@localhost:5432/postgres
export GOOSE_MIGRATION_DIR=./examples/goose/pg/ex1/migrations
export GOOSE_TABLE=public.goose_migrations

COMMAND="$1"

case $COMMAND in
    "down")
        goose down # Drop constrain
        sleep 2
        goose down # Drop columns
        sleep 2
        goose down # Drop tables
        ;;
    "status")
        goose status
        ;;
    "up")
        goose up
        ;;
    *)
        echo "Usage: $0 [down | up]"
        ;;
esac