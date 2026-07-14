package main

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/about"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool, tm *templates.TemplateManager) *chi.Mux {
	mx := chi.NewRouter()

	about.RegisterRoutes(mx, pool, tm)

	return mx
}
