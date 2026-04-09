package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// We declare 'db' here so all functions in this file can use it.
var db *sql.DB

// --- FUNCTION 1: The "Log Handler" ---
// This function runs every time someone visits http://localhost:8080/logs
func getLogsHandler(w http.ResponseWriter, r *http.Request) {
	// 'w' (ResponseWriter) is your "outgoing" megaphone to the browser.
	// 'r' (Request) is the "incoming" envelope from the browser.

	fmt.Fprintf(w, "Checking the database for your logs...")

	// Eventually, you will write a SQL query here using the 'db' variable
	// to fetch the 30-day trends and send them back as JSON.
}

// --- FUNCTION 2: The "Main" Engine ---
func main() {
	var err error
	connStr := "postgres://user_admin:secret_password@localhost:5432/health_tracker_db?sslmode=disable"

	// 1. Initialize the Database Connection
	// We use the global 'db' variable we declared at the top.
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Verify the Connection
	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("✅ Database is ready.")

	// 3. Routing (The "Switchboard")
	// This tells the Go server: "If someone asks for /logs, run the getLogsHandler function."
	http.HandleFunc("/logs", getLogsHandler)

	// 4. Start the Web Server
	// This line "blocks" the program. It stays here forever, listening on Port 8080.
	fmt.Println("🚀 Server starting on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
