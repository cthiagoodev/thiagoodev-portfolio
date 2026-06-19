package main

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/about"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
	mx := chi.NewRouter()

	mx.Route("/api/v1", func(router chi.Router) {
		about.RegisterRoutes(router, pool)
	})

	return mx
}
