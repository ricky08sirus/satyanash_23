# Go Webserver Project Documentation

## Project Overview
This project is a Go-based web server designed to interact with a PostgreSQL database. It provides endpoints for managing tables and their metadata, including creation, modification, and data insertion. The server is Dockerized for easy deployment.

### Features
- **Dynamic Table Creation**: Create tables with custom schemas.
- **Table Metadata Management**: Store and retrieve metadata related to tables.
- **Data Insertion**: Insert data into dynamically created tables.
- **PostgreSQL Integration**: Connects to a PostgreSQL database to manage tables and data.

---

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Setup

1. **Clone the Repository**:
   ```bash
   git clone git clone https://gitlab.com/ravij2/aman08jul.git


2. **Create the .env File: Create a .env file in the root directory with the following content**:


POSTGRES_USER=postgres
POSTGRES_PASSWORD=6394
POSTGRES_DB=table_data
DB_HOST=postgres_container

```bash
docker-compose up --build
```


# API Endpoints

## 1. Create Table

## Endpoint: 
POST /create-table

## Description: 
This endpoint creates a new table in the database based on the user-defined schema and also creates a corresponding metadata table. The metadata table stores information about the table's purpose, column descriptions, and creation timestamp. Additionally, a column column_name is added to the metadata table.

**Request Body** (JSON):

```json
{
  "table_name": "example_table",
  "table_purpose": "A brief description",
  "columns": [
    {
      "name": "column1",
      "type": "VARCHAR(255)",
      "description": "Description of column1"
    },
    {
      "name": "column2",
      "type": "INTEGER",
      "description": "Description of column2"
    }
  ]
}
```



## Success Response:

Status Code: 201 Created
Content-Type: application/json
Error Responses:

400 Bad Request: If the request body is missing required fields or contains invalid data.
500 Internal Server Error: If there is an error creating the table or metadata.


## 2. Insert Data

## Endpoint: POST /insert-data

### Description: 

This endpoint inserts data into a specified table and updates the metadata table with information about the insertion.
**Request Body** (JSON):
```json
{
  "table_name": "example_table",
  "data": {
    "column1": "value1",
    "column2": 123
  }
}
```

## Success Response:

Status Code: 
200 OK
Content-Type: 
application/json

## Error Responses:

400 Bad Request: If the request body is missing required fields or contains invalid data.
404 Not Found: If the specified table does not exist.
500 Internal Server Error: If there is an error inserting data or updating the metadata.


## 3. Get Metadata

## Endpoint: 
GET /metadata/{tableName}

## Description: 
Retrieve metadata for a specific table.

## Request Parameters:

## table_name (query parameter): 
The name of the table for which metadata is to be retrieved.

**Response Body** (JSON):
```json
{
  "table_name": "example_table",
  "created_at": "2024-09-12T10:45:00Z",
  "columns": [
    {
      "name": "column1",
      "description": "Description of column1"
    },
    {
      "name": "column2",
      "description": "Description of column2"
    }
  ]
}
```

## Success Response:

Status Code: 
200 OK
Content-Type: 
application/json



## 4. Modify Table

## Endpoint: 
POST /modify-table

## Description: 
Modify an existing table in the database by adding or updating columns. This also updates the corresponding metadata table.
**Request Body** (JSON):
   ```json
{
  "table_name": "example_table",
  "modifications": [
    {
      "action": "add_column",
      "column_name": "new_column",
      "column_type": "TEXT"
    }
  ]
}
```

## Success Response:

Status Code: 200 OK
Content-Type: application/json
## Error Responses:

400 Bad Request: If the request body is missing required fields or contains invalid data.
404 Not Found: If the specified table does not exist.
500 Internal Server Error: If there is an error modifying the table or updating metadata.












