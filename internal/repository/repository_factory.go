package repository

import (
	"context"

	"gyad/internal/database"
	"gyad/repositories"
	"gyad/models"
)

// RepositoryFactory is a factory for creating repositories
type RepositoryFactory struct {
	db *database.Database
}

// NewRepositoryFactory creates a new RepositoryFactory instance
func NewRepositoryFactory(db *database.Database) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

// * Your repository interfaces and implementations go here *
type BobersRepository interface {
	Create(ctx context.Context, bober *models.Bobers) error
	FindByID(ctx context.Context, id string) (*models.Bobers, error)
	Update(ctx context.Context, bober *models.Bobers) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Bobers, error)
}

func (f *RepositoryFactory) CreateBoberRepository() BobersRepository {
    return repositories.NewPgBobersRepository(f.db.Conn)
}