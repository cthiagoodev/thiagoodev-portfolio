package repositories

import (
	"context"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
)

type ProjectsRepository interface {
	GetAll(ctx context.Context) ([]entities.Project, error)
	ResetAndCreateAll(ctx context.Context, projects []entities.Project) error
}
