package gateways

import (
	"context"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
)

type ProjectsGateway interface {
	PublishProjects(ctx context.Context, projects []entities.Project) error
}
