package handlers

import (
	"fmt"
	"go-todo/cmd/storage"
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
