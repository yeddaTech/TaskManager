package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeddaTech/TaskManager/internals/handlers"
	"github.com/yeddaTech/TaskManager/internals/models" // Assicurati che ci sia
	"github.com/yeddaTech/TaskManager/templates"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Index().Render(r.Context(), w)
	})

	// DASHBOARD DINAMICA
	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		// Qui dovresti chiamare una funzione che recupera i task dal DB
		// Per ora passiamo una lista vuota per far compilare, poi userai handlers.GetTasks(r)
		var tasks []models.Task
		templates.Dashboard(tasks).Render(r.Context(), w)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		templates.Login().Render(r.Context(), w)
	})

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		templates.Register().Render(r.Context(), w)
	})

	// WORK DINAMICO
	r.Get("/work", func(w http.ResponseWriter, r *http.Request) {
		var t models.Task // Recupera il task attivo dal DB
		templates.Work(t).Render(r.Context(), w)
	})

	// PROFILO DINAMICO
	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		var u models.User // Recupera l'utente loggato dal DB
		templates.Profile(u).Render(r.Context(), w)
	})

	// GESTIONE AZIONI (POST)
	r.Post("/register", handlers.PostRegister)
	r.Post("/login", handlers.PostLogin)
	r.Post("/logout", handlers.PostLogout)
	r.Post("/tasks", handlers.PostTask) // Rotta per creare nuovi task

	return r
}
