package storage

import (
	models "go-todo/cmd/models"
)

type Storage interface {
	CreateTodo(*models.Todo) (int, error)
	DeleteTodo(string) error
	UpdateTodo(*models.Todo) error
	GetTodos() ([]*models.Todo, error)
	GetTodoByID(ID string) (*models.Todo, error)
}
