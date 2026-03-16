package api

import (
	"net/http"
	"sync"

	"github.com/yeddaTech/TaskManager/internals/db"
	"github.com/yeddaTech/TaskManager/internals/router"
)

var once sync.Once

func Handler(w http.ResponseWriter, r *http.Request) {
	// Inizializza il database solo la prima volta che la funzione viene svegliata
	once.Do(func() {
		db.InitDB()
	})

	// Usiamo il router condiviso
	myRouter := router.New()
	myRouter.ServeHTTP(w, r)
}
