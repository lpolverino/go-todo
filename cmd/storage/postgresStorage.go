package storage

import (
	"database/sql"
	"fmt"
	"go-todo/cmd/models"
	"log"
	"os"
	"strings"

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

func (s *PostgresStorage) CreateTodo(newTodo *models.Todo) (int, error) {
	var returningId int
	query := `insert into 
	todos (title, author, content, status)
	values ($1, $2, $3, $4)
	RETURNING id`

	rsp, err := s.db.Query(query,
		newTodo.Title, newTodo.Author, newTodo.Content, newTodo.Status,
	)

	if err != nil {
		return 0, err
	}

	rsp.Next()
	rsp.Scan(&returningId)

	return returningId, nil
}

func (s *PostgresStorage) DeleteTodo(id string) error {
	query := `Delete from todos where id = $1`
	_, err := s.db.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) UpdateTodo(newTodo *models.Todo) error {

	query := "UPDATE todos SET "
	var updates []string
	var args []interface{}

	argID := 1

	if newTodo.Title != "" {
		updates = append(updates, fmt.Sprintf("title = $%d", argID))
		args = append(args, newTodo.Title)
		argID++
	}
	if newTodo.Author != "" {
		updates = append(updates, fmt.Sprintf("author = $%d", argID))
		args = append(args, newTodo.Author)
		argID++
	}
	if newTodo.Content != "" {
		updates = append(updates, fmt.Sprintf("content = $%d", argID))
		args = append(args, newTodo.Content)
		argID++
	}

	updates = append(updates, fmt.Sprintf("status = $%d", argID))
	args = append(args, newTodo.Status)
	argID++

	query += strings.Join(updates, ", ")
	query += fmt.Sprintf(" WHERE id = $%d RETURNING id", argID)
	args = append(args, newTodo.Id)

	_, err := s.db.Query(query, args...)
	return err
}

func (s *PostgresStorage) GetTodos() ([]*models.Todo, error) {
	query := `select * from todos`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	todos := []*models.Todo{}

	for rows.Next() {
		todo := new(models.Todo)
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Author, &todo.Content, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func (s *PostgresStorage) GetTodoByID(ID string) (*models.Todo, error) {
	//TODO: input santizing is requierid
	query := `select * from todos where id = $1`
	rsp, err := s.db.Query(query, ID)
	if err != nil {
		return nil, err
	}
	if !rsp.Next() {
		return nil, nil
	}
	todo := new(models.Todo)
	err = rsp.Scan(&todo.Id, &todo.Title, &todo.Author, &todo.Content, &todo.Status)
	log.Printf("The todo is %+v", todo)
	return todo, err
}
