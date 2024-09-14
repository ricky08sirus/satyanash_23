package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ModifyTableRequest represents the incoming request structure
type ModifyTableRequest struct {
	TableName string `json:"table_name"`
	Operation string `json:"operation"`
	Column    struct {
		Name string `json:"name"`
		Type string `json:"type,omitempty"` // Type is required for add_column
	} `json:"column"`
}

// ModifyTableHandler handles table modifications
func ModifyTableHandler(w http.ResponseWriter, r *http.Request) {
	var req ModifyTableRequest

	// Parse the request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate operation
	if req.Operation != "add_column" && req.Operation != "remove_column" {
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	var query string

	// Build query based on operation
	switch req.Operation {
	case "add_column":
		if req.Column.Name == "" || req.Column.Type == "" {
			http.Error(w, "Column name and type are required for adding a column", http.StatusBadRequest)
			return
		}
		query = fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", req.TableName, req.Column.Name, req.Column.Type)
	case "remove_column":
		if req.Column.Name == "" {
			http.Error(w, "Column name is required for removing a column", http.StatusBadRequest)
			return
		}
		query = fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s;", req.TableName, req.Column.Name)
	}

	// Execute the query
	_, err = db.Exec(query)
	if err != nil {
		log.Printf("Failed to modify table: %v", err)
		http.Error(w, "Failed to modify table", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Table modified successfully"))
}
