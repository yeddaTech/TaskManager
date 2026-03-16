package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	Pool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic("Impossibile connettersi al database: " + err.Error())
	}
}
