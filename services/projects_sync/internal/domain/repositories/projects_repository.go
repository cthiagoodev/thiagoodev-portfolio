package repositories

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type ProjectsRepository interface {
	GetAll() ([]entities.Project, error)
	CreateOrUpdate(project entities.Project) error
	Delete(uuid string) error
}
