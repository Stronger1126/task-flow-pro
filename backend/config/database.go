package config

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

// SetupDB initializes the database connection
func SetupDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    return db, nil
} 

// DataSource creates a data source name for connecting to PostgreSQL
func DataSource(user, password, dbname, host string, port int) string {
    return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", user, password, dbname, host, port)
}