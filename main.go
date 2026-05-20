package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Local fallback for running outside Docker
		dbURL = "postgres://devops_user:super_secret_password@localhost:5432/sandbox_analytics?sslmode=disable"
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(ctx)

	fmt.Println("Go API successfully connected to PostgreSQL!")

	// Demo endpoint to show system metric logs
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var serviceName string
		// To ensure connection works
		err := conn.QueryRow(ctx, "SELECT service_name FROM system_logs LIMIT 1").Scan(&serviceName)
		if err != nil {
			// Can be empty on a fresh run so lets echo a success link
			fmt.Fprintf(w, `{"data":{"status":"Go API connected. No metric logs written yet."}}`)
			return
		}

		fmt.Fprintf(w, `{"data":{"status":"Go API functional", "latest_log_from":"%s"}}`, serviceName)
	})

	log.Println("Go Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
