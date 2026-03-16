package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeddaTech/TaskManager/internals/handlers"
	"github.com/yeddaTech/TaskManager/templates"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	// Helper per controllare se l'utente è loggato (per il Layout)
	isLogged := func(r *http.Request) bool {
		cookie, err := r.Cookie("user_id")
		return err == nil && cookie.Value != ""
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Contiamo gli utenti
		count := handlers.GetUserCount()
		// Passiamo sia il login che il conteggio al template!
		templates.Index(isLogged(r), count).Render(r.Context(), w)
	})

	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		// Recuperiamo i task reali dal database tramite l'handler
		tasks := handlers.GetTasksFromDB(r)
		templates.Dashboard(tasks, isLogged(r)).Render(r.Context(), w)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		if isLogged(r) {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
		templates.Login(isLogged(r)).Render(r.Context(), w)
	})

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		templates.Register(isLogged(r)).Render(r.Context(), w)
	})

	r.Get("/work", func(w http.ResponseWriter, r *http.Request) {
		task := handlers.GetActiveTaskFromDB(r)
		templates.Work(task, isLogged(r)).Render(r.Context(), w)
	})

	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		user := handlers.GetUserFromDB(r)
		templates.Profile(user, isLogged(r)).Render(r.Context(), w)
	})

	r.Get("/tasks/new", func(w http.ResponseWriter, r *http.Request) {
		templates.NewTask(isLogged(r)).Render(r.Context(), w)
	})

	// DA COSÌ:
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//     templates.Index(isLogged(r)).Render(r.Context(), w)
	// })

	// A COSÌ:

	// GESTIONE AZIONI (POST)
	r.Post("/register", handlers.PostRegister)
	r.Post("/login", handlers.PostLogin)
	r.Post("/logout", handlers.PostLogout)
	r.Post("/tasks", handlers.PostTask)
	r.Post("/tasks/complete/{id}", handlers.PostCompleteTask)
	r.Post("/tasks/start/{id}", handlers.PostStartTask)

	return r
}
