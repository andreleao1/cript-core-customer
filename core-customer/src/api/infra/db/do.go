package db

import (
	"log/slog"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

const dbUrl = "postgres://postgres:postgres@localhost:5432/core-customer?sslmode=disable"

func OpenConnection() *sqlx.DB {
	slog.Info("Opening connection to database")

	connection, err := sqlx.Open("pgx", dbUrl)

	if err != nil {
		slog.Error("Error opening connection to database: %v", err)
		panic(err)
	}

	if err = connection.Ping(); err != nil {
		slog.Error("Error pinging database: %v", err)
		panic(err)
	}

	slog.Info("Connected to database")
	return connection
}
