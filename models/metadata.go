package models

import "time"

type TableMetadata struct {
    TableName   string    `json:"table_name"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    Columns     []Column  `json:"columns"`
}

