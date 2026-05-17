package main

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/server/internal/about"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool) *chi.Mux {
	router := chi.NewRouter()

	about.RegisterRoutes(router, pool)

	return router
}
