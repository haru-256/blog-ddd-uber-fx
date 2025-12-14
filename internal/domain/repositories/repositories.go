package repositories

import (
	"context"

	"github.com/haru-256/blog-ddd-uber-fx/internal/domain/models"
)

//go:generate go tool mockgen -source=$GOFILE -destination=../../mocks/repository/repository_mock.go -package=mock_repository
type UserRepository interface {
	GetAll(ctx context.Context) ([]*models.User, error)
	GetById(ctx context.Context, id *models.UserId) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id *models.UserId) error
}

//go:generate go tool mockgen -source=$GOFILE -destination=../../mocks/repository/repository_mock.go -package=mock_repository
type TaskRepository interface {
	GetAll(ctx context.Context) ([]*models.Task, error)
	GetById(ctx context.Context, id *models.TaskId) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) (*models.Task, error)
	Update(ctx context.Context, task *models.Task) (*models.Task, error)
	Delete(ctx context.Context, id *models.TaskId) error
}
