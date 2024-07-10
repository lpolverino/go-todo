package handlers

import (
	"fmt"
	"go-todo/cmd/models"
	"go-todo/cmd/storage"
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context, s storage.Storage) error {
	todos, err := s.GetTodos()
	if err != nil {
		return fmt.Errorf("cannot get todos, %+v", err)
	}
	return c.JSON(http.StatusOK, todos)
}

func GetTodo(c echo.Context, s storage.Storage) error {
	todo, err := s.GetTodoByID(c.Param("todoId"))
	if err != nil {
		return fmt.Errorf("cannot get todo, %+v", err)
	}
	return c.JSON(http.StatusOK, todo)
}

func AddTodo(c echo.Context, s storage.Storage) error {
	newTodo := models.Todo{}
	err := c.Bind(&newTodo)

	if err != nil {
		log.Printf("There was an error with the body binding %+v", err)
		return c.String(http.StatusBadRequest, "bad body")
	}
	todoId, err := s.CreateTodo(&newTodo)

	if err != nil {
		return fmt.Errorf("cannot add todo, %+v", err)
	}

	return c.String(http.StatusCreated, fmt.Sprintf("The Todo was created %d", todoId))
}

func UpdateTodo(c echo.Context, s storage.Storage) error {
	newValueForTodo := models.Todo{}
	err := c.Bind(&newValueForTodo)

	if err != nil {
		log.Printf("There was an erro with the body binding %+v", err)
		return c.String(http.StatusBadRequest, "bad body")
	}
	//TODO: santize input
	newValueForTodo.Id = c.Param("todoId")
	err = s.UpdateTodo(&newValueForTodo)
	if err != nil {
		return fmt.Errorf("there was an error updating the Todo, %+v", err)
	}

	return c.String(http.StatusOK, "The todo was modified")
}

func DeleteTodo(c echo.Context, s storage.Storage) error {
	//TODO: sanitizr input
	deleteId := c.Param("todoId")
	err := s.DeleteTodo(deleteId)
	if err != nil {
		return fmt.Errorf("there was an erro updating the Todo, %+v", err)
	}
	return c.String(http.StatusOK, "The todo was deleted")
}
