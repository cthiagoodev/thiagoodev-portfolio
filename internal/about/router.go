package about

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *chi.Mux, pool *pgxpool.Pool) {
	const path = "/about"

	repository := NewDatabaseRepository(pool)
	useCase := NewUseCase(repository)
	handler := NewHandler(useCase)

	router.Get(path, handler.Get)
}
