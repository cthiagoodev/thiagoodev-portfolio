package about

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router chi.Router, pool *pgxpool.Pool, tm *templates.TemplateManager) {
	const path = "/about"

	repository := NewDatabaseRepository(pool)
	useCase := NewUseCase(repository)
	handler := NewHandler(useCase, tm)

	router.Get(path, handler.Get)
}
