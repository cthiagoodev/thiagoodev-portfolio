package usecases

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type CreateOrUpdateProjectUseCase interface {
	Execute(project entities.Project) error
}
