package service

import (
	"context"
	"crud/models"
)

// EmpRepo
type EmpRepo interface {
	Fetch(ctx context.Context) ([]*models.Employee, error)
	GetByID(ctx context.Context, id int64) (*models.Employee, error)
	Create(ctx context.Context, e *models.Employee) (int64, error)
	Update(ctx context.Context, e *models.Employee, id int64) (*models.Employee, error)
	Delete(ctx context.Context, id int64) (int64, error)
	Trasaction(ctx context.Context, amount float64, senderId int64, receiverId int64) ( error)
}

// UserRepository
type UserRepo interface {
	Login(ctx context.Context, username string, password string) (*string, error)
	Register(ctx context.Context, u *models.User) (int64, error)
	Fetch(ctx context.Context) ([]*models.User, error)
}
