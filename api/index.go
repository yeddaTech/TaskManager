package handler

import (
	"net/http"

	"github.com/yeddaTech/TaskManager/cmd/prod"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	prod.MyWebApp(w, r)
}
