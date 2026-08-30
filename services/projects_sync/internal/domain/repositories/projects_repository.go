package repositories

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type ProjectsRepository interface {
	CreateAll(projects []entities.Project) ([]entities.Project, error)
	DeleteAll() error
}
