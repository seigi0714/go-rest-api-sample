package main

import (
	"fmt"

	"github.com/seigi0714/go-rest-api-sample/internal/router"
)

func main() {
	fmt.Print("serve")
	router.Handle()
}
