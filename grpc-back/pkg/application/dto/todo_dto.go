package dto

import "github.com/jumpei00/grpc/grpcback/pkg/domain"

type Todo struct {
	Key      string
	Content  string
}

func NewTodo(todo *domain.Todo) *Todo {
	return &Todo{
		Key:      todo.Key,
		Content:  todo.Content,
	}
}

func NewTodos(todos []*domain.Todo) []*Todo {
	ts := make([]*Todo, len(todos))
	for i, t := range todos {
		ts[i] = &Todo{
			Key:      t.Key,
			Content:  t.Content,
		}
	}
	return ts
}
