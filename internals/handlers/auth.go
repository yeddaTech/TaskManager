package handlers

import (
	"context"
	"net/http"

	"github.com/yeddaTech/TaskManager/internals/db"
	"golang.org/x/crypto/bcrypt"
)

// Registrazione
func PostRegister(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
		username, email, string(hashedPassword))

	if err != nil {
		http.Error(w, "Errore nella registrazione", http.StatusInternalServerError)
		return
	}
	// Dopo la registrazione lo mandiamo al login
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Login semplice con Cookie
func PostLogin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	var id int
	var hash string
	err := db.Pool.QueryRow(context.Background(),
		"SELECT id, password_hash FROM users WHERE email = $1", email).Scan(&id, &hash)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		http.Error(w, "Credenziali errate", http.StatusUnauthorized)
		return
	}

	// Settiamo un cookie ignorante con l'ID utente (per ora va bene così)
	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: string(rune(id)),
		Path:  "/",
	})
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Logout
func PostLogout(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{

		Name: "user_id",

		Value: "",

		Path: "/",

		MaxAge: -1, // Questo dice al browser di cancellare il cookie immediatamente

		HttpOnly: true,
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
