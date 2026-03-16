package handlers

import (
	"context"
	"net/http"
	"strconv" // <-- AGGIUNTO PER CONVERTIRE I NUMERI CORRETTAMENTE

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

// Login
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

	// COOKIE BLINDATO PER VERCEL E CONVERSIONE ID CORRETTA
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    strconv.Itoa(id), // <-- CORRETTO! Ora trasforma l'ID 1 nel testo "1"
		Path:     "/",
		HttpOnly: true,                 // Non accessibile da JavaScript
		Secure:   true,                 // OBBLIGATORIO SU VERCEL (HTTPS)
		SameSite: http.SameSiteLaxMode, // Permette il login senza blocchi strani
		MaxAge:   86400 * 7,            // Dura una settimana
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Logout
func PostLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // Uccide il cookie all'istante
		HttpOnly: true,
		Secure:   true, // Allineato con il cookie di login
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Aggiungi questo in fondo al file auth.go
func GetUserCount() int {
	var count int
	// Chiediamo al database di contare tutte le righe nella tabella users
	err := db.Pool.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return 0 // Se c'è un errore, mostriamo 0 per non far crashare nulla
	}
	return count
}
