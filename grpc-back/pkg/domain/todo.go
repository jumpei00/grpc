package domain

import "github.com/google/uuid"

type Todo struct {
	Key      string
	Content  string
}

func NewTodo(content string) *Todo {
	uuid := uuid.NewString()
	return &Todo{
		Key:      uuid,
		Content:  content,
	}
}

func (t *Todo) Update(content string) {
	t.Content = content
}
