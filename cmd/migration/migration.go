package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run migration.go create <migration_name> | migrate <all/migration_name>")
		os.Exit(1)
	}

	command := os.Args[1]
	migrationName := os.Args[2]

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	switch command {
		case "create":
			createMigration(migrationName)
		case "migrate":
			migrate(migrationName)
		default:
			fmt.Println("Invalid command. Use 'create' or 'migrate'.")
			os.Exit(1)
	}
}

// createMigration creates a new migration file with the given name.
func createMigration(migrationName string) {
	timestamp := time.Now().Format("20060102150405")
	migrationsDir := "./migrations"
	filename := fmt.Sprintf("%s/%s_%s.sql", migrationsDir, timestamp, migrationName)

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		if err := os.MkdirAll(migrationsDir, os.ModePerm); err != nil {
			fmt.Printf("Failed to create migrations directory: %v\n", err)
			os.Exit(1)
		}
	}

	content := `-- Up migration
-- Write SQL statements for applying the migration here.

-- Down migration
-- Write SQL statements for reverting the migration here.
`
	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		fmt.Printf("Failed to create migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Migration created at: %s\n", filename)
}

// migrate applies all or specific migrations to the database.
func migrate(argument string) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSL")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSL)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	files, err := ioutil.ReadDir("./migrations")
	if err != nil {
		log.Fatalf("Could not read migrations directory: %v", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if argument == "all" || strings.Contains(file.Name(), argument) {
			fmt.Printf("Applying migration: %s\n", file.Name())
			err := executeSQLFile(db, filepath.Join("./migrations", file.Name()))
			if err != nil {
				log.Fatalf("Failed to apply migration %s: %v", file.Name(), err)
			}
		}
	}
}

// executeSQLFile reads the SQL file and executes the statements in it.
func executeSQLFile(db *sql.DB, filepath string) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	requests := strings.Split(string(content), "-- Down migration")
	sqlStatements := strings.Split(requests[0], ";")

	for _, statement := range sqlStatements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}
		_, err := db.Exec(statement)
		if err != nil {
			return err
		}
	}
	return nil
}
