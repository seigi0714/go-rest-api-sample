package router

import (
	"github.com/gorilla/mux"
	"github.com/seigi0714/go-rest-api-sample/pkg/controller"
)

var tr controller.TodoRouter = controller.NewTodoRouter()

func Handle() {
	myRouter := mux.NewRouter().StrictSlash(false)
	myRouter.HandleFunc("/todos", tr.HandleTodosRequest)
}
