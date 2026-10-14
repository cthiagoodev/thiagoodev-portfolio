package usecases

import (
	"context"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
)

type GetAboutUseCase interface {
	Execute(ctx context.Context) (*entities.About, error)
}
