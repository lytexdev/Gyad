package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database represents a database connection
type Database struct {
	Conn *sql.DB
}

// NewDatabase creates a new Database instance
func NewDatabase() *Database {
	return &Database{}
}

// Connect connects to the database
func (db *Database) Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbSSL)

	var err error
	db.Conn, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = db.Conn.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to database")
}

// Close closes the database connection
func (db *Database) Close() {
	if err := db.Conn.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	fmt.Println("Database connection closed... mom")
}

// Query executes a query that returns rows, typically a SELECT statement
func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Exec executes a query without returning any rows
func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
