package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/tomasbasham/blunderlist-todo/internal"
	"github.com/tomasbasham/blunderlist-todo/internal/grpc"
	"github.com/tomasbasham/blunderlist-todo/internal/storage/pg"
)

func main() {
	logger := log.New(os.Stdout, "", log.Lshortfile)

	connectionString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("BLUNDERLIST_TODO_GCLOUD_SQLPROXY_SERVICE_HOST"),
		os.Getenv("BLUNDERLIST_TODO_GCLOUD_SQLPROXY_SERVICE_PORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Fatalf("unable to connect to the database: %s", err)
	}

	defer db.Close()

	// Setup dependencies for the service.
	storage := pg.NewStorage(db)
	service := internal.NewStore(storage)

	todoService := grpc.NewServer(logger, service)
	todoServicePort := os.Getenv("BLUNDERLIST_TODO_SERVICE_PORT")

	if err := todoService.Serve(todoServicePort); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
