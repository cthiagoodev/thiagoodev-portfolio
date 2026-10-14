package repositories

import (
	"context"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
)

type AboutRepository interface {
	Get(ctx context.Context) (*entities.About, error)
}
