package about

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router chi.Router, pool *pgxpool.Pool) {
	repository := NewDatabaseRepository(pool)
	useCase := NewUseCase(repository)
	handler := NewHandler(useCase)

	router.Get("/about", handler.Get)
}
