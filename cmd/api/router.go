package main

import (
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/about"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/home"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool, tm *templates.TemplateManager) http.Handler {
	router := chi.NewRouter()

	home.RegisterRoutes(router, tm)
	about.RegisterRoutes(router, pool, tm)

	return router
}
