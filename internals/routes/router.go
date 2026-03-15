package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeddaTech/TaskManager/templates"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		templates.Index().Render(r.Context(), w)
	})

	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		templates.Dashboard().Render(r.Context(), w)
	})

	// ROTTA 3: Login
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Login().Render(r.Context(), w)
	})

	// ROTTA 4: Register
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Register().Render(r.Context(), w)
	})

	// ROTTA 5: Work
	r.Get("/work", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Work().Render(r.Context(), w)
	})

	// ROTTA 6: Profile
	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Profile().Render(r.Context(), w)
	})

	return r
}
