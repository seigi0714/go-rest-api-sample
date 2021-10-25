package main

import (
	"net/http"

	"github.com/seigi0714/go-rest-api-sample/internal/router"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	router.Handle()
	server.ListenAndServe()
}
