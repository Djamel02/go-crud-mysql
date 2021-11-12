package service

import (
	"context"
	"crud/models"
)

// EmpRepo explain...
type EmpRepo interface {
	Fetch(ctx context.Context) ([]*models.Employee, error)
	GetByID(ctx context.Context, id int64) (*models.Employee, error)
	Create(ctx context.Context, e *models.Employee) (int64, error)
	Update(ctx context.Context, e *models.Employee, id int64) (*models.Employee, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
