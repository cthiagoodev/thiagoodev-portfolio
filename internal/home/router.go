package home

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, tm *templates.TemplateManager) {
	const path = "/"

	handler := NewHandler(tm)

	router.Get(path, handler.Get)
}
