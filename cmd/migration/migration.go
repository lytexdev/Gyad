package main

import (
	"database/sql"
	"fmt"
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
		fmt.Println("Usage: go run cmd/migration/migration.go create <migration_name> | migrate <all/migration_name> | rollback <migration_name>")
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
	case "rollback":
		rollbackMigration(migrationName)
	default:
		fmt.Println("Invalid command. Use 'create', 'migrate', or 'rollback'.")
		os.Exit(1)
	}
}

func createMigration(migrationName string) {
	migrationsDir := "./migrations"
	timestamp := time.Now().Format("20060102150405")
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

func migrate(argument string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	files, err := os.ReadDir("./migrations")
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

func rollbackMigration(argument string) {
    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

    files, err := os.ReadDir("./migrations")
    if err != nil {
        log.Fatalf("Could not read migrations directory: %v", err)
    }

    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

    for _, file := range files {
        if strings.Contains(file.Name(), argument) {
            fmt.Printf("Rolling back migration: %s\n", file.Name())
            err := executeSQLDown(db, filepath.Join("./migrations", file.Name()))
            if err != nil {
                log.Fatalf("Failed to rollback migration %s: %v", file.Name(), err)
            }

            if askForConfirmation(fmt.Sprintf("Do you want to delete the migration file %s?", file.Name())) {
                err := os.Remove(filepath.Join("./migrations", file.Name()))
                if err != nil {
                    log.Printf("Failed to delete migration file: %v", err)
                } else {
                    fmt.Printf("Migration file %s deleted successfully.\n", file.Name())
                }
            }
        }
    }
}

func executeSQLFile(db *sql.DB, filepath string) error {
	content, err := os.ReadFile(filepath)
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

func executeSQLDown(db *sql.DB, filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	requests := strings.Split(string(content), "-- Down migration")
	sqlStatements := strings.Split(requests[1], ";")

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

func askForConfirmation(question string) bool {
    fmt.Println(question + " [y/n]: ")
	
    var response string

    _, err := fmt.Scanln(&response)
    if err != nil {
        log.Printf("Invalid input: %v", err)
        return false
    }
    response = strings.TrimSpace(response)
    return strings.ToLower(response) == "y" || strings.ToLower(response) == "yes"
}
