package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "go-rest-api-sample/pkg/controller/todo"
)

var tr controller.TodoRouter = controller.NewRouter()

func Handle() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/todos", tr.HandleTodosRequest)
	myRouter.HandleFunc("/todos/{id}", tr.HandleTodosRequest)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: myRouter,
	}
	log.Fatal(server.ListenAndServe())
}
