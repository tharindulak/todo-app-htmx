package todo

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type ITodoService interface {
	GetAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	Create(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(id int) error
}

type TodoService struct {
	db *sql.DB
}

func (s *TodoService) GetAll() ([]Todo, error) {
	rows, err := s.db.Query("SELECT * FROM todos")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

func (s *TodoService) GetByID(id int) (Todo, error) {
	var todo Todo

	row := s.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *TodoService) Create(todo Todo) (Todo, error) {
	_, err := s.db.Exec(
		"INSERT INTO todos (title, description, status) VALUES (?, ?, ?)",
		todo.Title,
		todo.Description,
		todo.Status)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *TodoService) Update(todo Todo) (Todo, error) {
	_, err := s.db.Exec(
		"UPDATE todos SET title = ?, description = ?, status = ? WHERE id = ?",
		todo.Title,
		todo.Description,
		todo.Status,
		todo.ID)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *TodoService) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func NewTodoService(db *sql.DB) ITodoService {
	return &TodoService{db}
}

func RegisterTodoHandler(r *gin.Engine) {

}
