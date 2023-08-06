package application

import (
	"context"

	"github.com/jumpei00/grpc/grpcback/pkg/application/dto"
	"github.com/jumpei00/grpc/grpcback/pkg/domain"
	"github.com/jumpei00/grpc/grpcback/pkg/domain/repository"
)

type TodoApplication interface {
	GetAll(ctx context.Context) ([]*dto.Todo, error)
	Get(ctx context.Context, key string) (*dto.Todo, error)
	Create(ctx context.Context, content string) (*dto.Todo, error)
	Update(ctx context.Context, key, content string) (*dto.Todo, error)
	Delete(ctx context.Context, key string) error
}

type todoApplication struct {
	todoRepos repository.TodoRepository
}

func NewTodoApplication(todoRepos repository.TodoRepository) TodoApplication {
	return &todoApplication{
		todoRepos: todoRepos,
	}
}

func (t *todoApplication) GetAll(ctx context.Context) ([]*dto.Todo, error) {
	todos, err := t.todoRepos.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return dto.NewTodos(todos), nil
}

func (t *todoApplication) Get(ctx context.Context, key string) (*dto.Todo, error) {
	todo, err := t.todoRepos.GetByKey(ctx, key)
	if err != nil {
		return nil, err
	}
	return dto.NewTodo(todo), nil
}

func (t *todoApplication) Create(ctx context.Context, content string) (*dto.Todo, error) {
	newTodo := domain.NewTodo(content)

	todo, err := t.todoRepos.Create(ctx, newTodo)
	if err != nil {
		return nil, err
	}

	return dto.NewTodo(todo), nil
}

func (t *todoApplication) Update(ctx context.Context, key, content string) (*dto.Todo, error) {
	todo, err := t.todoRepos.GetByKey(ctx, key)
	if err != nil {
		return nil, err
	}

	todo.Update(content)

	updatedTodo, err := t.todoRepos.Update(ctx, todo)
	if err != nil {
		return nil, err
	}

	return dto.NewTodo(updatedTodo), nil
}

func (t *todoApplication) Delete(ctx context.Context, key string) error {
	return t.todoRepos.Delete(ctx, key)
}
