package todo_controller

import (
	"fmt"
	"go-rest-api-sample/pkg/model/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoRouter interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

type todoRouter struct {
	tc TodoController
}

func NewRouter() TodoRouter {
	var tr = repository.NewTodoRepository()
	var tc = NewTodoController(tr)
	return &todoRouter{tc}
}

func (ro *todoRouter) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]

	fmt.Printf("method :: %s", r.Method)
	switch r.Method {
	case "GET":
		if todoId != "" {
			ro.tc.GetTodo(w, r)
		} else {
			ro.tc.GetTodos(w, r)
		}
	case "POST":
		ro.tc.PostTodo(w, r)
	case "PUT":
		ro.tc.PutTodo(w, r)
	case "DELETE":
		ro.tc.DeleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}
