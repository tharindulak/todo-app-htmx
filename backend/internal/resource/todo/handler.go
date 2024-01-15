package todo

import "database/sql"

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

func (s *TodoService) GetAll() ([]Todo, error) {}

func (s *TodoService) GetByID(id int) (Todo, error) {}

func (s *TodoService) Create(todo Todo) (Todo, error) {}

func (s *TodoService) Update(todo Todo) (Todo, error) {}

func (s *TodoService) Delete(id int) error {}

func NewTodoService(db *sql.DB) ITodoService {
	return &TodoService{db}
}
