package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/database"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/modules"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatalf("failed to initialize database url")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx := context.Background()

	pool, err := database.NewPool(ctx, databaseURL)
	if err != nil {
		log.Fatalf("failed to initialize database pool: %v", err)
	}
	defer pool.Close()

	applicationModule := modules.NewModule(pool)
	chiRouter := router.NewRouter(applicationModule)

	log.Printf("server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, chiRouter); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
