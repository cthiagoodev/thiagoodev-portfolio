package main

import (
	"io/fs"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/about"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/home"
	templates2 "github.com/cthiagoodev/thiagoodev-portfolio/templates"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(pool *pgxpool.Pool, tm *templates.TemplateManager) (http.Handler, error) {
	router := chi.NewRouter()

	if err := NewStaticRouter(router); err != nil {
		return nil, err
	}

	home.RegisterRoutes(router, tm)
	about.RegisterRoutes(router, pool, tm)

	return router, nil
}

func NewStaticRouter(router chi.Router) error {
	staticFS, err := fs.Sub(templates2.StaticsFS, "static")

	if err != nil {
		return err
	}

	fileServer := http.FileServer(http.FS(staticFS))
	staticHandler := http.StripPrefix("/static/", fileServer)
	router.Handle("/static/*", staticHandler)

	return nil
}
