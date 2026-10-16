package usecases

import (
	"context"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
)

type GetAllSocialContactsUseCase interface {
	Execute(ctx context.Context) ([]entities.SocialContact, error)
}

