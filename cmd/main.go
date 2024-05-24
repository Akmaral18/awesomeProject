package main

import (
	"awesomeProject"
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

const dbsource = "postgresql://akmaral:1qazXSW@@localhost:5433/postgres?sslmode=disable"

func main() {
	var httpAddr = flag.String("http", ":8000", "http listen address")
	logger := slog.Default()

	ctx := context.Background()

	logger.InfoContext(ctx, "initializing data connection...")
	db, err := sql.Open("postgres", dbsource)
	if err != nil {
		logger.ErrorContext(ctx, "failed to initialize a connection to postgres", "err", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		logger.ErrorContext(ctx, "failed to ping data", "err", err)
		os.Exit(1)
	}

	logger.InfoContext(ctx, "data connection successfully initialized, building routes")

	serviceContainer := awesomeProject.NewServiceContainer(logger, db)
	router := awesomeProject.NewRouter(logger, serviceContainer)

	logger.InfoContext(ctx, fmt.Sprint("routes successfully initialized, now listening on port 8000"))

	if err := http.ListenAndServe(*httpAddr, router); err != nil {
		logger.ErrorContext(ctx, "failed to start server", "err", err)
		os.Exit(1)
	}
}
