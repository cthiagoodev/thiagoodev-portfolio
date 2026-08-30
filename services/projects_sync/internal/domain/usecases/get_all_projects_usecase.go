package usecases

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type GetAllProjectsUseCase interface {
	Execute() ([]entities.Project, error)
}
