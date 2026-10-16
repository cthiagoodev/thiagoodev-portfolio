package repositories

import (
	"context"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
)

type SocialContactRepository interface {
	GetAll(ctx context.Context) ([]entities.SocialContact, error)
}

