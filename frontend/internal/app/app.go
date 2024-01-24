package app

import (
	"encoding/json"
	"net/http"

	common "github.com/charukak/todo-app-htmx/common/pkg"
)

// API client for todo app
// /todos GET
// /todos/{id} GET
// /todos POST
// /todos/{id} PUT
// /todos/{id} DELETE
type TodoAppClient struct {
	url string
}

// Create a new TodoAppClient
func NewTodoAppClient(url string) *TodoAppClient {
	return &TodoAppClient{url}
}

// Get all todos
func (c *TodoAppClient) GetTodos() ([]common.Todo, error) {
	req, err := http.NewRequest("GET", c.url+"/todos", nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var todos []common.Todo

	err = json.NewDecoder(resp.Body).Decode(&todos)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Get a todo by id
func (c *TodoAppClient) GetTodoById(id string) (*common.Todo, error) {
	req, err := http.NewRequest("GET", c.url+"/todos/"+id, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var todo common.Todo

	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

// Create a new todo
func (c *TodoAppClient) CreateTodo() (*common.Todo, error) {
	req, err := http.NewRequest("POST", c.url+"/todos", nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var todo common.Todo

	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (c *TodoAppClient) UpdateTodoById(id string) (*common.Todo, error) {
	req, err := http.NewRequest("PUT", c.url+"/todos/"+id, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var todo common.Todo

	err = json.NewDecoder(resp.Body).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (c *TodoAppClient) DeleteTodoById(id string) error {
	req, err := http.NewRequest("DELETE", c.url+"/todos/"+id, nil)

	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
