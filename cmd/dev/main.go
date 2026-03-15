package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yeddaTech/TaskManager/templates"
)

func pageHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	templates.Index().Render(r.Context(), w)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", pageHandle)
	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Dashboard().Render(r.Context(), w)
	})

	// NUOVE ROTTE:
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Login().Render(r.Context(), w)
	})

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Register().Render(r.Context(), w)
	})

	r.Get("/work", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Work().Render(r.Context(), w)
	})

	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		templates.Profile().Render(r.Context(), w)
	})

	fs := http.FileServer(http.Dir("public/dist"))
	r.Handle("/dist/*", http.StripPrefix("/dist/", fs))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
