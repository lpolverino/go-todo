package storage

import (
	models "go-todo/cmd/models"
)

type Storage interface {
	CreateTodo(*models.Todo) error
	DeleteTodo(int) error
	UpdateTodo(*models.Todo) error
	GetTodos() ([]*models.Todo, error)
	GetTodoByID(ID int) (*models.Todo, error)
}
