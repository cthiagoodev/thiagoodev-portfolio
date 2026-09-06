package repositories

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type ProjectsRepository interface {
	GetAll() ([]entities.Project, error)
	CreateAll(projects []entities.Project) error
	DeleteAll() error
}
