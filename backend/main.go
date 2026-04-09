package main

import (
	"database/sql"
	"fmt"
	"log"

	// This driver acts as the bridge between Go's standard SQL package and Postgres
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// 1. Define your connection string
	// Format: postgres://username:password@host:port/database_name
	connStr := "postgres://user_admin:secret_password@localhost:5432/mood_tracker_db"

	// 2. Open a connection pool (this doesn't actually connect yet)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// 3. Ensure the connection is closed when the program finishes
	defer db.Close()

	// 4. The Handshake: Actually try to communicate with the DB
	fmt.Println("Attempting to connect to the database...")

	// We use a small loop because sometimes the DB takes a second to wake up
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database handshake failed: %v", err)
	}

	fmt.Println("✅ Successfully connected to the database!")
}

///////// HTTP Server and Handlers would go here, but we're focusing on the database connection for this snippet.
