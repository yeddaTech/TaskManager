package main

import (
	"fmt"
	"net/http"

	"github.com/yeddaTech/TaskManager/internals/router"
)

func main() {
	r := router.New()
	fmt.Println("Server locale partito su http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
