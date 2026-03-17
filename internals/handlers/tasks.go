package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/yeddaTech/TaskManager/internals/db"
	"github.com/yeddaTech/TaskManager/internals/models"
)

// Crea un nuovo task
func PostTask(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user_id")
	title := r.FormValue("title")
	desc := r.FormValue("description")
	deadlineStr := r.FormValue("deadline")

	deadline, _ := time.Parse("2006-01-02", deadlineStr)

	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO tasks (user_id, title, description, deadline, status) VALUES ($1, $2, $3, $4, 'pending')",
		cookie.Value, title, desc, deadline)

	if err != nil {
		http.Error(w, "Errore inserimento task", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Completa un task (Tasto FINE)
func PostCompleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	_, err := db.Pool.Exec(context.Background(), "UPDATE tasks SET status = 'completed' WHERE id = $1", taskID)
	if err != nil {
		http.Error(w, "Errore update", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Prende i task per la Dashboard (e fa pulizia automatica)
func GetTasksFromDB(r *http.Request) []models.Task {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		return nil
	}

	// 1. IL BOIA: Elimina fisicamente i task completati la cui scadenza è passata
	// CURRENT_DATE è oggi. "< CURRENT_DATE" significa da ieri in giù.
	deleteQuery := `
        DELETE FROM tasks 
        WHERE user_id = $1 
        AND status = 'completed' 
        AND deadline < CURRENT_DATE
    `
	_, _ = db.Pool.Exec(r.Context(), deleteQuery, cookie.Value)

	// 2. SELEZIONE GENUINA: Ora peschiamo esattamente quello che è rimasto, senza filtri strani
	selectQuery := `
        SELECT id, title, description, status, deadline 
        FROM tasks 
        WHERE user_id = $1
    `
	rows, err := db.Pool.Query(r.Context(), selectQuery, cookie.Value)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Deadline)
		tasks = append(tasks, t)
	}
	return tasks
}

// Prende il task attivo per la pagina Work
func GetActiveTaskFromDB(r *http.Request) models.Task {
	cookie, err := r.Cookie("user_id")
	var t models.Task
	if err != nil {
		return t
	}

	db.Pool.QueryRow(r.Context(), "SELECT id, title, deadline FROM tasks WHERE user_id = $1 AND status != 'completed' LIMIT 1", cookie.Value).Scan(&t.ID, &t.Title, &t.Deadline)
	return t
}

// Prende l'utente per il Profilo
func GetUserFromDB(r *http.Request) models.User {
	cookie, err := r.Cookie("user_id")
	var u models.User
	if err != nil {
		return u
	}
	db.Pool.QueryRow(r.Context(), "SELECT id, username, email FROM users WHERE id = $1", cookie.Value).Scan(&u.ID, &u.Username, &u.Email)
	return u
}
func PostStartTask(w http.ResponseWriter, r *http.Request) {
	// 1. Prendiamo l'ID del task dall'URL (es: /tasks/start/5)
	taskID := chi.URLParam(r, "id")

	// 2. Diciamo al database di aggiornare lo stato
	_, err := db.Pool.Exec(context.Background(),
		"UPDATE tasks SET status = 'doing' WHERE id = $1",
		taskID,
	)

	if err != nil {
		http.Error(w, "Errore nell'avvio del task", http.StatusInternalServerError)
		return
	}

	// 3. Ricarichiamo la pagina della Dashboard per vedere la magia
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
