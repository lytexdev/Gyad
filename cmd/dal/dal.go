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

    ensureDirectories([]string{"models", "repository"})

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

        fmt.Println("Enter data type (int, string, float, bool, or foreign key like User:ID): ")
        dataType, _ := reader.ReadString('\n')
        dataType = strings.TrimSpace(dataType)

        foreignKey := ""
        if strings.Contains(dataType, ":") {
            parts := strings.Split(dataType, ":")
            dataType = fmt.Sprintf("*%s", parts[0])
            foreignKey = parts[1]
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

type {{.EntityName}} struct {
    ID   string ` + "`db:\"id\"`" + `
    {{- range .Fields}}
    {{.Name}} {{.Type}} ` + "`db:\"{{.Name}}\"{{if eq .IsNullable \"yes\"}} nullable{{end}}`" + ` // Field
    {{- if .ForeignKey}}
    // Foreign key relation
    {{.Name}}ID {{.Type}} ` + "`db:\"{{.ForeignKey}}\"`" + `
    {{- end}}
    {{- end}}
}

func New{{.EntityName}}() *{{.EntityName}} {
    return &{{.EntityName}}{}
}
`
    tmpl, err := template.New("model").Parse(modelTemplate)
    if err != nil {
        return err
    }

    modelFile, err := os.Create(fmt.Sprintf("models/%s.go", strings.ToLower(entityName)))
    if err != nil {
        return err
    }
    defer modelFile.Close()

    return tmpl.Execute(modelFile, map[string]interface{}{
        "EntityName": entityName,
        "Fields":     fields,
    })
}

func generateRepository(entityName, tableName string, fields []Field) error {
    repoTemplate := `package repository

import (
    "database/sql"
    "log"
    "gyad/models"
)

type {{.EntityName}}Repository struct {
    db *sql.DB
}

func New{{.EntityName}}Repository(db *sql.DB) *{{.EntityName}}Repository {
    return &{{.EntityName}}Repository{db: db}
}

func (r *{{.EntityName}}Repository) Create(m *models.{{.EntityName}}) error {
    // Insert SQL code here
    return nil
}

func (r *{{.EntityName}}Repository) Update(m *models.{{.EntityName}}) error {
    // Update SQL code here
    return nil
}

func (r *{{.EntityName}}Repository) Delete(id string) error {
    // Delete SQL code here
    return nil
}

func (r *{{.EntityName}}Repository) GetAll() ([]*models.{{.EntityName}}, error) {
    // GetAll SQL code here
    return nil
}

func (r *{{.EntityName}}Repository) GetByID(id string) (*models.{{.EntityName}}, error) {
    // GetByID SQL code here
    return nil
}
`
    tmpl, err := template.New("repository").Parse(repoTemplate)
    if err != nil {
        return err
    }

    repoFile, err := os.Create(fmt.Sprintf("repository/%s_repository.go", strings.ToLower(entityName)))
    if err != nil {
        return err
    }
    defer repoFile.Close()

    return tmpl.Execute(repoFile, map[string]interface{}{
        "EntityName": entityName,
        "TableName":  tableName,
    })
}
