package main

import (
	"net/http"

	"github.com/seigi0714/go-rest-api-sample/controller"
	"github.com/seigi0714/go-rest-api-sample/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
