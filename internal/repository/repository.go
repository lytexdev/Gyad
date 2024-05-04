package repository

import (
	"context"
	"gyad/models"
)

type BoberRepository interface {
	Create(ctx context.Context, bober *models.Bobers) error
	FindByID(ctx context.Context, id string) (*models.Bobers, error)
	Update(ctx context.Context, bober *models.Bobers) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Bobers, error)
}

type MamRepository interface {
	Create(ctx context.Context, bober *models.Bobers) error
	FindByID(ctx context.Context, id string) (*models.Bobers, error)
	Update(ctx context.Context, bober *models.Bobers) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]models.Bobers, error)
}
