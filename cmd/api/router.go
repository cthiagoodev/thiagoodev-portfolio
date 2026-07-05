package main

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/about"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
	mx := chi.NewRouter()

	about.RegisterRoutes(mx, pool)

	return mx
}
