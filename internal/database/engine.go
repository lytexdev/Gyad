package database

import (
    "fmt"
    "log"
    "os"

    "xorm.io/xorm"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

type Database struct {
    Engine *xorm.Engine
}

// NewEngine creates a new database connection
func NewEngine() (*Database, error) {
    if err := godotenv.Load(); err != nil {
        log.Printf("Error loading .env file: %v", err)
        return nil, err
    }

    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

    engine, err := xorm.NewEngine("postgres", connectionString)
    if err != nil {
        log.Printf("Failed to create an XORM engine: %v", err)
        return nil, err
    }

    if err = engine.Ping(); err != nil {
        log.Printf("Failed to connect to the database: %v", err)
        return nil, err
    }

    fmt.Println("Connected to database using XORM")
    return &Database{Engine: engine}, nil
}

// Close closes the database connection
func (db *Database) Close() {
    if err := db.Engine.Close(); err != nil {
        log.Printf("Error closing database connection: %v", err)
    } else {
        fmt.Println("Database connection closed")
    }
}
