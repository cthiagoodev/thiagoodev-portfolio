package supabase

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, os.Getenv("SUPABASE_DATABASE_URL"))
}
