package about

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(mx *chi.Mux, pool *pgxpool.Pool) {
	repository := NewDatabaseRepository(pool)
	useCase := NewUseCase(repository)
	handler := NewHandler(useCase)

	mx.Get("/v1/about", handler.Get)
}
