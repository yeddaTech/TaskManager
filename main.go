package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moroz/oikaze/templates"
)

func pageHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	templates.Index().Render(r.Context(), w)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", pageHandle)

	fs := http.Dir("./assets/dist")
	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(fs)))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
