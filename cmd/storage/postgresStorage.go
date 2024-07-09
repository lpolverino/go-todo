package storage

import (
	"database/sql"
	"fmt"
	"go-todo/cmd/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbPort))

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) InitDB() error {
	if err := s.CreateUsersTable(); err != nil {
		return err
	}
	if err := s.CreateTodosTable(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) CreateTodosTable() error {
	query := `create table if not exists todos (
	id serial primary key,
		title varchar(80),
		author varchar(50),
		content varchar(1000),
		status boolean
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateUsersTable() error {
	return nil
}

func (s *PostgresStorage) CreateTodo(*models.Todo) error {
	return nil
}
func (s *PostgresStorage) DeleteTodo(int) error {
	return nil
}
func (s *PostgresStorage) UpdateTodo(*models.Todo) error {
	return nil
}
func (s *PostgresStorage) GetTodos() ([]*models.Todo, error) {
	return nil, nil
}
func (s *PostgresStorage) GetTodoByID(ID int) (*models.Todo, error) {
	return nil, nil
}
