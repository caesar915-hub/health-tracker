package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// We declare 'db' here so all functions in this file can use it.
var db *sql.DB

// --- FUNCTION 1: The "Get Logs" Handler ---
// This function will run when someone visits http://localhost:8080/logs
// http.ResponseWriter is how we send data back to the browser
// *http.Request is how we get information about the incoming request (like URL, headers, etc.)
func getLogsHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Write the SQL Query
	// This selects logs from the last 30 days, ordered by date
	query := `
		SELECT id, entry_date, sleep_quality, physical_energy, focus, 
		       motivation, past_view, social_activity, created_at 
		FROM daily_metrics 
		WHERE entry_date > CURRENT_DATE - INTERVAL '30 days'
		ORDER BY entry_date ASC`

	// 2. Execute the query
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close() // to close the connection when we're done

	// 3. Create a Slice to hold our results
	var logs []DailyMetrics

	// 4. Iterate through the rows
	for rows.Next() {
		var m DailyMetrics
		// The order in Scan must match the order in your SELECT statement!
		// We scan the current row into our DailyMetrics struct
		// scan uses pointers (&) to fill the struct fields with data from the database
		err := rows.Scan(
			&m.ID, &m.EntryDate, &m.SleepQuality, &m.PhysicalEnergy,
			&m.Focus, &m.Motivation, &m.PastView, &m.SocialActivity, &m.CreatedAt,
		)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}
		logs = append(logs, m)
	}

	// 5. Convert the Slice to JSON and send it
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// --- FUNCTION 2: The "Create Log" Handler ---
// This function will run when someone sends a POST request to http://localhost:8080/logs
func createLogHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Safety Check: Only allow "POST" requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Create an empty instance of your Struct
	var entry DailyMetrics

	// 3. "Decode" the JSON from the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// 4. The SQL Insert
	// This query uses placeholders ($1, $2, etc.) to safely insert data without risking SQL injection
	query := `
		INSERT INTO daily_metrics (
			entry_date, sleep_quality, physical_energy, 
			focus, motivation, past_view, social_activity
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Exec(query,
		entry.EntryDate, entry.SleepQuality, entry.PhysicalEnergy,
		entry.Focus, entry.Motivation, entry.PastView, entry.SocialActivity)

	if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	// 5. Success Message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Log saved successfully!")
}

// --- FUNCTION 3: The "Main" Engine ---
func main() {
	var err error
	connStr := "postgres://user_admin:secret_password@localhost:5432/health_tracker_db?sslmode=disable"

	// 1. Initialize the Database Connection - GO METHOD 1: sql.Open
	// We use the global 'db' variable we declared at the top.
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Verify the Connection - GO METHOD 1: db.Ping()
	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("✅ Database is ready.")

	// 3. Routing (The "Switchboard") - GO METHOD 1: http.HandleFunc
	// This tells the Go server: "If someone asks for /logs, run the getLogsHandler function."
	http.HandleFunc("/logs", getLogsHandler)
	http.HandleFunc("/submit", createLogHandler)

	// 1. Create a "File Server" handler that points to your frontend folder
	fs := http.FileServer(http.Dir("../frontend"))

	// 2. Tell the router: "For the root path (/), use that File Server"
	http.Handle("/", fs)

	// 4. Start the Web Server GO METHOD 1: http.ListenAndServe
	// This line "blocks" the program. It stays here forever, listening on Port 8080.
	fmt.Println("🚀 Server starting on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
