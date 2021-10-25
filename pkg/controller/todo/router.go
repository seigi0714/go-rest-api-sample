package controller

import (
	"net/http"

	"github.com/seigi0714/go-rest-api-sample/model/repository"
)

type TodoRouter interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

type todoRouter struct {
	tc TodoController
}

func NewTodoRouter() TodoRouter {
	var tr = repository.NewTodoRepository()
	var tc = NewTodoController(tr)
	return NewRouter(tc)
}

func NewRouter(tc TodoController) TodoRouter {
	return &todoRouter{tc}
}

func (ro *todoRouter) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetTodos(w, r)
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
