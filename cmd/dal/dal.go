package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "text/template"
)

type Field struct {
    Name         string
    Type         string
    IsNullable   string
    ForeignKey   string
}

func main() {
    if len(os.Args) < 2 || os.Args[1] != "create" {
        fmt.Println("Usage: ./dal create")
        os.Exit(1)
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Entity Name in PascalCase: ")
    entityName, _ := reader.ReadString('\n')
    entityName = strings.TrimSpace(entityName)

    fmt.Print("Enter Table Name in snake_case: ")
    tableName, _ := reader.ReadString('\n')
    tableName = strings.TrimSpace(tableName)

    fields := collectFields(reader)

    ensureDirectories([]string{"models", "repositories"})

    if err := generateModel(entityName, tableName, fields); err != nil {
        fmt.Printf("Error generating model: %v\n", err)
        os.Exit(1)
    }

    if err := generateRepository(entityName, tableName, fields); err != nil {
        fmt.Printf("Error generating repository: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("DAL files created successfully.")
}

func collectFields(reader *bufio.Reader) []Field {
    var fields []Field
    for {
        fmt.Print("Enter field name (or press Enter to finish): ")
        fieldName, _ := reader.ReadString('\n')
        fieldName = strings.TrimSpace(fieldName)
        if fieldName == "" {
            break
        }

        fmt.Print("Enter data type (int, string, float, bool, date etc.): ")
        dataType, _ := reader.ReadString('\n')
        dataType = strings.TrimSpace(dataType)

        var foreignKey string
        fmt.Print("Is this field a foreign key? (Enter referenced model and field, e.g., User:ID, or type no): ")
        fkInput, _ := reader.ReadString('\n')
        fkInput = strings.TrimSpace(fkInput)
        if fkInput != "no" {
            foreignKey = fkInput
        }

        fmt.Print("Can the field be nullable? (yes/no): ")
        nullable, _ := reader.ReadString('\n')
        nullable = strings.TrimSpace(nullable)

        fields = append(fields, Field{Name: fieldName, Type: dataType, IsNullable: nullable, ForeignKey: foreignKey})
    }
    return fields
}

func ensureDirectories(dirs []string) {
    for _, dir := range dirs {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            os.MkdirAll(dir, 0755)
        }
    }
}

func generateModel(entityName, tableName string, fields []Field) error {
    modelTemplate := `package models

import "time"

// {{.EntityName}} represents the database model for the table "{{.TableName}}"
type {{.EntityName}} struct {
    ID        string    ` + "`db:\"id\"`" + ` // Primary key
    CreatedAt time.Time ` + "`db:\"created_at\"`" + ` // Timestamp for when this record was created
    UpdatedAt time.Time ` + "`db:\"updated_at\"`" + ` // Timestamp for when this record was last updated
    {{- range .Fields}}
    {{.Name}} {{if eq .Type "time.Time"}}*{{end}}{{.Type}} ` + "`db:\"{{lower .Name}}\"{{if eq .IsNullable \"yes\"}},omitempty{{end}}`" + ` // Field
    {{if .ForeignKey}}
    {{.Name}}ID string ` + "`db:\"{{lower .ForeignKey}}\"`" + ` // Foreign key to {{.ForeignKey}}
    {{end}}
    {{- end}}
}
`
    tmpl, err := template.New("model").Funcs(template.FuncMap{
        "lower": strings.ToLower,
    }).Parse(modelTemplate)
    if err != nil {
        return err
    }

    modelFile, err := os.Create(fmt.Sprintf("models/%s_model.go", strings.ToLower(entityName)))
    if err != nil {
        return err
    }
    defer modelFile.Close()

    return tmpl.Execute(modelFile, map[string]interface{}{
        "EntityName": entityName,
        "TableName": tableName,
        "Fields":     fields,
    })
}

func generateRepository(entityName, tableName string, fields []Field) error {
    repoTemplate := `package repositories

import (
    "context"
    "database/sql"
    "gyad/models"
)

// {{.EntityName}}Repository provides methods to perform CRUD operations on {{.EntityName}} entities.
type {{.EntityName}}Repository struct {
    db *sql.DB
}

// New{{.EntityName}}Repository creates a new instance of {{.EntityName}}Repository
func New{{.EntityName}}Repository(db *sql.DB) *{{.EntityName}}Repository {
    return &{{.EntityName}}Repository{db: db}
}

// Create inserts a new {{.EntityName}} into the database, setting ID, created_at, and updated_at automatically
func (r *{{.EntityName}}Repository) Create(ctx context.Context, m *models.{{.EntityName}}) error {
    query := "INSERT INTO {{.TableName}} ({{range .Fields}}{{.Name}}, {{end}}created_at, updated_at) VALUES ({{range .Fields}}?, {{end}}NOW(), NOW())"
    _, err := r.db.ExecContext(ctx, query, {{range .Fields}}m.{{.Name}}, {{end}})
    return err
}

// Update modifies an existing {{.EntityName}} in the database
func (r *{{.EntityName}}Repository) Update(ctx context.Context, m *models.{{.EntityName}}) error {
    query := "UPDATE {{.TableName}} SET {{range .Fields}}{{.Name}} = ?, {{end}}updated_at = NOW() WHERE id = ?"
    _, err := r.db.ExecContext(ctx, query, {{range .Fields}}m.{{.Name}}, {{end}}m.ID)
    return err
}

// Delete removes a {{.EntityName}} from the database
func (r *{{.EntityName}}Repository) Delete(ctx context.Context, id string) error {
    query := "DELETE FROM {{.TableName}} WHERE id = ?"
    _, err := r.db.ExecContext(ctx, query, id)
    return err
}

// GetAll retrieves all {{.EntityName}} entities from the database
func (r *{{.EntityName}}Repository) GetAll(ctx context.Context) ([]*models.{{.EntityName}}, error) {
    query := "SELECT * FROM {{.TableName}}"
    rows, err := r.db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []*models.{{.EntityName}}
    for rows.Next() {
        var m models.{{.EntityName}}
        if err := rows.Scan(&m); err != nil {
            return nil, err
        }
        results = append(results, &m)
    }
    return results, nil
}

// GetByID retrieves a specific {{.EntityName}} by ID from the database
func (r *{{.EntityName}}Repository) GetByID(ctx context.Context, id string) (*models.{{.EntityName}}, error) {
    query := "SELECT * FROM {{.TableName}} WHERE id = ?"
    var m models.{{.EntityName}}
    err := r.db.QueryRowContext(ctx, query, id).Scan(&m)
    if err != nil {
        return nil, err
    }
    return &m, nil
}
`
    tmpl, err := template.New("repository").Funcs(template.FuncMap{
        "lower": strings.ToLower,
    }).Parse(repoTemplate)
    if err != nil {
        return err
    }

    repoFile, err := os.Create(fmt.Sprintf("repositories/%s_repository.go", strings.ToLower(entityName)))
    if err != nil {
        return err
    }
    defer repoFile.Close()

    return tmpl.Execute(repoFile, map[string]interface{}{
        "EntityName": entityName,
        "TableName":  tableName,
        "Fields":     fields,
    })
}
