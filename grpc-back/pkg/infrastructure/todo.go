package infrastructure

import (
	"context"
	"database/sql"

	"github.com/jumpei00/grpc/grpcback/pkg/domain"
	"github.com/jumpei00/grpc/grpcback/pkg/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) GetAll(ctx context.Context) ([]*domain.Todo, error) {
	todoModels, err := models.Todos().All(ctx, t.db)
	if err != nil {
		return nil, err
	}

	var todos []*domain.Todo
	for _, todoModel := range todoModels {
		todos = append(todos, &domain.Todo{
			Key:     todoModel.Key,
			Content: todoModel.Content.String,
		})
	}

	return todos, nil
}

func (t *todoRepository) GetByKey(ctx context.Context, key string) (*domain.Todo, error) {
	todoModel, err := models.FindTodo(ctx, t.db, key)
	if err != nil {
		return nil, err
	}

	return &domain.Todo{
		Key:     todoModel.Key,
		Content: todoModel.Content.String,
	}, nil
}

func (t *todoRepository) Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	todoModel := &models.Todo{
		Key:     todo.Key,
		Content: null.StringFrom(todo.Content),
	}

	if err := todoModel.Insert(ctx, tx, boil.Infer()); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return todo, nil
}

func (t *todoRepository) Update(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	todoModel, err := models.FindTodo(ctx, tx, todo.Key)
	if err != nil {
		return nil, err
	}

	todoModel.Content = null.StringFrom(todo.Content)

	if _, err := todoModel.Update(ctx, tx, boil.Infer()); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	return todo, nil
}

func (t *todoRepository) Delete(ctx context.Context, key string) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	todoModel, err := models.FindTodo(ctx, tx, key)
	if err != nil {
		return err
	}

	if _, err := todoModel.Delete(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}
