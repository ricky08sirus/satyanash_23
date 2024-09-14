package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// InsertDataHandler handles data insertion into the main table and metadata updates
func InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		TableName string                 `json:"table_name"`
		Data      map[string]interface{} `json:"data"`
	}

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare the INSERT statement for the main table
	columns := ""
	placeholders := ""
	values := []interface{}{}
	i := 1

	for col, val := range requestData.Data {
		if columns != "" {
			columns += ", "
			placeholders += ", "
		}
		columns += col
		placeholders += fmt.Sprintf("$%d", i)
		values = append(values, val)
		i++
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", requestData.TableName, columns, placeholders)

	// Execute the INSERT query for the main table
	_, err := db.Exec(query, values...)
	if err != nil {
		log.Printf("Failed to insert data: %v", err)
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	// Insert metadata into the metadata table (ensure all entries are stored)
	err = insertMetadata(requestData.TableName, requestData.Data)
	if err != nil {
		log.Printf("Failed to insert metadata: %v", err)
		http.Error(w, "Failed to insert metadata", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data inserted and metadata updated successfully"))
}

// insertMetadata inserts metadata into the metadata table
func insertMetadata(tableName string, data map[string]interface{}) error {
	// Construct the metadata table name dynamically
	metadataTable := tableName + "_meta_data"

	// Iterate over columns and insert metadata for each column
	for col, val := range data {
		query := `INSERT INTO ` + metadataTable + ` (table_name, column_name, created_at, description) 
                  VALUES ($1, $2, CURRENT_TIMESTAMP, $3);`

		_, err := db.Exec(query, tableName, col, "Column "+col+" added with value "+fmt.Sprint(val))
		if err != nil {
			log.Printf("Error inserting into metadata table %s: %v", metadataTable, err)
			return err
		}
	}
	return nil
}
