package todo_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"go-rest-api-sample/pkg/controller/dto"
	"go-rest-api-sample/pkg/model/entity"
	"go-rest-api-sample/pkg/model/repository"
)

type TodoController interface {
	GetTodo(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	PutTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repository.TodoRepository
}

func NewTodoController(tr repository.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}
	//TODO :: 取得するFieldをクエリパラメータで指定
	todo, err := tc.tr.GetTodo(todoId, []string{"id"})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	output, _ := json.MarshalIndent(todo, "", "\t\t")
	fmt.Println("get Todo: ", output)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *todoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	// TODO:: 取得するFieldをクエリパラメータで指定
	todos, err := tc.tr.GetTodos([]string{"id", "content"})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todosResponse dto.TodosResponse
	todosResponse.Todos = todos

	output, _ := json.MarshalIndent(todos, "", "\t\t")
	fmt.Println("get Todo: ", output)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)
	fmt.Println("request ::", todoRequest)
	todo := entity.TodoEntity{Title: todoRequest.Title, Content: todoRequest.Content}
	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *todoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := entity.TodoEntity{Id: todoId, Title: todoRequest.Title, Content: todoRequest.Content}
	err = tc.tr.UpdateTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = tc.tr.DeleteTodo(todoId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
