package about

import (
	"github.com/cthiagoodev/portfolio/services/app-api/internal/application/usecases"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/repositories"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AboutModule struct {
	Handler *handlers.AboutHandler
}

func NewAboutModule(pool *pgxpool.Pool) *AboutModule {
	aboutRepository := repositories.NewAboutDatabaseRepository(pool)
	getAboutUseCase := usecases.NewGetAboutUseCase(aboutRepository)
	aboutHandler := handlers.NewAboutHandler(getAboutUseCase)

	return &AboutModule{
		Handler: aboutHandler,
	}
}
