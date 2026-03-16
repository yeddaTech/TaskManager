package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/yeddaTech/TaskManager/internals/db"
)

func PostTask(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user_id")
	title := r.FormValue("title")
	desc := r.FormValue("description")
	deadlineStr := r.FormValue("deadline") // Formato input date: 2026-03-25

	deadline, _ := time.Parse("2006-01-02", deadlineStr)

	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO tasks (user_id, title, description, deadline) VALUES ($1, $2, $3, $4)",
		cookie.Value, title, desc, deadline)

	if err != nil {
		http.Error(w, "Errore inserimento task", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
