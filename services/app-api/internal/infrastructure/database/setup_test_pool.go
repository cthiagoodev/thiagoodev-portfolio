//go:build integration

package database

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func SetupTestDatabase(t *testing.T) (context.Context, *pgxpool.Pool) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	_ = godotenv.Load("../../../.env")
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgresql://portfolio:change_me@localhost:5432/portfolio?sslmode=disable"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)

	pool, err := NewPool(ctx, databaseURL)
	if err != nil {
		t.Skipf("database is not reachable: %v", err)
	}

	t.Cleanup(func() {
		pool.Close()
	})

	return ctx, pool
}
