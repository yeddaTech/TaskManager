package api

import (
	"net/http"

	"github.com/yeddaTech/TaskManager/internals/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Usiamo il router condiviso
	myRouter := router.New()
	myRouter.ServeHTTP(w, r)
}
