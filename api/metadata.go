package api

import (
	//"database/sql" // Import the sql package for error handling
	"encoding/json"
	"net/http"
	"webserver/models" // Import the models package
)

// GetMetadataHandler handles requests to fetch table metadata
func GetMetadataHandler(w http.ResponseWriter, r *http.Request) {
	tableName := r.URL.Query().Get("table_name")
	if tableName == "" {
		http.Error(w, "Missing table_name parameter", http.StatusBadRequest)
		return
	}

	metadataList, err := fetchTableMetadata(tableName)
	if err != nil {
		http.Error(w, "Failed to fetch metadata", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(metadataList)
}

// fetchTableMetadata queries the database for metadata of the given table name
func fetchTableMetadata(tableName string) ([]models.TableMetadata, error) {
	var metadataList []models.TableMetadata

	// Construct the metadata table name dynamically
	metadataTable := tableName + "_meta_data"

	// Construct the query string dynamically
	query := `SELECT table_name, created_at, description FROM ` + metadataTable

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var meta models.TableMetadata
		err := rows.Scan(&meta.TableName, &meta.CreatedAt, &meta.Description)
		if err != nil {
			return nil, err
		}
		metadataList = append(metadataList, meta)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return metadataList, nil
}
