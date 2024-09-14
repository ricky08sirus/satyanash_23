package models

type Table struct {
    Name    string   `json:"name"`
    Columns []Column `json:"columns"`
}

