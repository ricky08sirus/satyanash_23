package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"webserver/api"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // Import the godotenv package
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	// Load the database URL from the .env file
	connStr := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
}

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	initDB()

	// Pass the db instance to API handlers
	api.SetDB(db)

	// Set up the router and routes
	r := mux.NewRouter()
	r.HandleFunc("/create-table", api.CreateTableHandler).Methods("POST")
	r.HandleFunc("/modify-table", api.ModifyTableHandler).Methods("POST")
	r.HandleFunc("/metadata", api.GetMetadataHandler).Methods("GET")
	r.HandleFunc("/insert-data", api.InsertDataHandler).Methods("POST")
	http.Handle("/", r)
	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
