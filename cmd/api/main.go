package main

import (
	"net/http"

	"github/seigi0714/go-rest-api-samle/controller"
	"github/seigi0714/go-rest-api-samle/model/repository"
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
