package repository

import (
	"context"

	"github.com/jumpei00/grpc/grpcback/pkg/domain"
)

type TodoRepository interface {
	GetAll(ctx context.Context) ([]*domain.Todo, error)
	GetByKey(ctx context.Context, key string) (*domain.Todo, error)
	Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)
	Delete(ctx context.Context, key string) error
}
