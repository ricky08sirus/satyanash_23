package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// CreateTableHandler handles creating the main table and metadata table
func CreateTableHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		TableName    string `json:"table_name"`
		TablePurpose string `json:"table_purpose"`
		Columns      []struct {
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"columns"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Create main table query
	mainTableQuery := "CREATE TABLE " + requestData.TableName + " ("
	for i, column := range requestData.Columns {
		if i > 0 {
			mainTableQuery += ", "
		}
		mainTableQuery += column.Name + " " + column.Type
	}
	mainTableQuery += ");"

	// Create metadata table query
	metaTableQuery := "CREATE TABLE " + requestData.TableName + "_meta_data ("
	metaTableQuery += "table_name VARCHAR(255) NOT NULL, "
	metaTableQuery += "description TEXT NOT NULL, "
	metaTableQuery += "column_name VARCHAR(255) NOT NULL, "
	metaTableQuery += "created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP"
	metaTableQuery += ");"

	// Execute queries
	_, err := db.Exec(mainTableQuery)
	if err != nil {
		log.Printf("Error executing main table query: %v", err)
		http.Error(w, "Failed to create main table", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(metaTableQuery)
	if err != nil {
		log.Printf("Error executing metadata table query: %v", err)
		http.Error(w, "Failed to create metadata table", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Tables created successfully"))
}
