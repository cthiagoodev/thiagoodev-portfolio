package social_contact

import (
	"github.com/cthiagoodev/portfolio/services/app-api/internal/application/usecases"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/infrastructure/repositories"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SocialContactModule struct {
	Handler *handlers.SocialContactHandler
}

func NewSocialContactModule(pool *pgxpool.Pool) *SocialContactModule {
	socialContactRepository := repositories.NewSocialContactDatabaseRepository(pool)
	getAllSocialContactsUseCase := usecases.NewGetAllSocialContactsUseCase(socialContactRepository)
	socialContactHandler := handlers.NewSocialContactHandler(getAllSocialContactsUseCase)

	return &SocialContactModule{
		Handler: socialContactHandler,
	}
}

