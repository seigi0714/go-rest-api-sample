package main

import (
	"fmt"

	"go-rest-api-sample/internal/router"
)

func main() {
	fmt.Print("serve")
	router.Handle()
}
