package repositories

import (
	"context"
	"database/sql"
	"gyad/models"
)

type PgBobersRepository struct {
	db *sql.DB
}

func NewPgBobersRepository(db *sql.DB) *PgBobersRepository {
	return &PgBobersRepository{db: db}
}

func (repo *PgBobersRepository) Create(ctx context.Context, bobers *models.Bobers) error {
	query := `INSERT INTO bobers (id, name, age, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := repo.db.ExecContext(ctx, query, bobers.ID, bobers.Name, bobers.Age, bobers.CreatedAt, bobers.UpdatedAt)
	return err
}

func (repo *PgBobersRepository) FindByID(ctx context.Context, id string) (*models.Bobers, error) {
	query := `SELECT id, name, age, created_at, updated_at FROM bobers WHERE id = $1`
	bobers := &models.Bobers{}
	err := repo.db.QueryRowContext(ctx, query, id).Scan(&bobers.ID, &bobers.Name, &bobers.Age, &bobers.CreatedAt, &bobers.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return bobers, nil
}

func (repo *PgBobersRepository) Update(ctx context.Context, bobers *models.Bobers) error {
	query := `UPDATE bobers SET name = $1, age = $2, updated_at = $3 WHERE id = $4`
	_, err := repo.db.ExecContext(ctx, query, bobers.Name, bobers.Age, bobers.UpdatedAt, bobers.ID)
	return err
}

func (repo *PgBobersRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM bobers WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}

func (repo *PgBobersRepository) List(ctx context.Context) ([]models.Bobers, error) {
	query := `SELECT id, name, age, created_at, updated_at FROM bobers`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bobers := []models.Bobers{}
	for rows.Next() {
		b := models.Bobers{}
		err = rows.Scan(&b.ID, &b.Name, &b.Age, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bobers = append(bobers, b)
	}
	return bobers, nil
}
