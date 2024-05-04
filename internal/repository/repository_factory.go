package repository

import (
	"gyad/internal/database"
	"gyad/repositories"
)

type RepositoryFactory struct {
	db *database.Database
}

// NewRepositoryFactory creates a new RepositoryFactory instance
func NewRepositoryFactory(db *database.Database) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) CreateBoberRepository() BoberRepository {
    return repositories.NewPgBobersRepository(f.db.Conn)
}