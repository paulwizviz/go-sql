package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	createTblPersonSQL = `CREATE TABLE IF NOT EXISTS person (
    uuid UUID PRIMARY KEY DEFAULT uuidv7()
);`
	createTblCollectivesSQL = `CREATE TABLE IF NOT EXISTS collectives (
    uuid UUID PRIMARY KEY DEFAULT uuidv7(),
    name TEXT NOT NULL
);`
	createTblMembersSQL = `CREATE TABLE IF NOT EXISTS collective_members (
    collective_uuid UUID NOT NULL,
    person_uuid UUID NOT NULL,
    PRIMARY KEY (collective_uuid, person_uuid),
    CONSTRAINT fk_collective FOREIGN KEY (collective_uuid) REFERENCES collectives(uuid) ON DELETE CASCADE,
    CONSTRAINT fk_person FOREIGN KEY (person_uuid) REFERENCES person(uuid) ON DELETE CASCADE
);`

	insertPersonSQL      = `INSERT INTO person DEFAULT VALUES RETURNING uuid`
	insertCollectiveSQL  = `INSERT INTO collectives (name) VALUES ($1) RETURNING uuid`
	addMemberSQL         = `INSERT INTO collective_members (collective_uuid, person_uuid) VALUES ($1, $2)`
	selectMembersSQL     = `SELECT person_uuid FROM collective_members WHERE collective_uuid = $1`
	dropTblMembersSQL    = `DROP TABLE IF EXISTS collective_members;`
	dropTblCollectivesSQL = `DROP TABLE IF EXISTS collectives;`
	dropTblPersonSQL      = `DROP TABLE IF EXISTS person;`
)

func main() {
	ctx := context.Background()

	// 1. Establish connection pool
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/default")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	// 2. Create Tables
	for _, sql := range []string{createTblPersonSQL, createTblCollectivesSQL, createTblMembersSQL} {
		if _, err := pool.Exec(ctx, sql); err != nil {
			log.Fatalf("Create table error: %v", err)
		}
	}

	// 3. Create multiple persons
	var personUUIDs []string
	for i := 0; i < 3; i++ {
		var u string
		if err := pool.QueryRow(ctx, insertPersonSQL).Scan(&u); err != nil {
			log.Fatalf("Insert person error: %v", err)
		}
		personUUIDs = append(personUUIDs, u)
		fmt.Printf("Created Person %d: %s\n", i+1, u)
	}

	// 4. Create a collective
	var collectiveUUID string
	if err := pool.QueryRow(ctx, insertCollectiveSQL, "Engineering Team").Scan(&collectiveUUID); err != nil {
		log.Fatalf("Insert collective error: %v", err)
	}
	fmt.Printf("Created Collective: %s\n", collectiveUUID)

	// 5. Link multiple persons to the collective
	for _, pUUID := range personUUIDs {
		if _, err := pool.Exec(ctx, addMemberSQL, collectiveUUID, pUUID); err != nil {
			log.Fatalf("Add member error: %v", err)
		}
	}

	// 6. Query and print members
	rows, err := pool.Query(ctx, selectMembersSQL, collectiveUUID)
	if err != nil {
		log.Fatalf("Select members error: %v", err)
	}
	defer rows.Close()

	fmt.Printf("\nMembers of '%s':\n", "Engineering Team")
	for rows.Next() {
		var u string
		if err := rows.Scan(&u); err != nil {
			log.Fatalf("Scan member error: %v", err)
		}
		fmt.Printf(" - %s\n", u)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Rows error: %v", err)
	}

	// 7. Cleanup
	for _, sql := range []string{dropTblMembersSQL, dropTblCollectivesSQL, dropTblPersonSQL} {
		if _, err := pool.Exec(ctx, sql); err != nil {
			log.Fatalf("Cleanup error: %v", err)
		}
	}
}
