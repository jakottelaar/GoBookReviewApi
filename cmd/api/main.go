package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jakottelaar/gobookreviewapp/api"
	"github.com/jakottelaar/gobookreviewapp/config"
	_ "github.com/jakottelaar/gobookreviewapp/docs"
	"github.com/jakottelaar/gobookreviewapp/pkg/database"
)

// @title			Book Review API
// @version		1.0
// @description	This is a sample server Book Review server.
// @host			localhost:8080
// @BasePath		/v1/api
// @schemes		http
// @produces		json
// @consumes		json
func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	log.Printf("Environment: %s", cfg.Environment)

	err = database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer database.Close()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	router := api.SetupRoutes()

	logger.Info("Starting server", "port", cfg.Port, "Environment", cfg.Environment)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
