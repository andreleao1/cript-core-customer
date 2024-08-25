package main

import (
	"core-customer/api/infra/db"
	"core-customer/api/routers"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	customLog := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(customLog)

	db := db.OpenConnection()

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		slog.Error("Error creating driver.")
		panic(err)
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	if err != nil {
		slog.Error("Error creating migrations.")
		panic(err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		slog.Error("Error running migrations.")
		panic(err)
	}

	router := gin.Default()
	contextPath := router.Group("/core-customer")
	{
		routers.Init(contextPath, db)
	}

	router.Run(":9092")
}
