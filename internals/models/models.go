package models

import (
	"time"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
}

type Task struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Status      string
	Deadline    time.Time // Questo gestisce la tua scadenza
	CreatedAt   time.Time
}
