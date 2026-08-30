package repositories

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type ProjectsDatabaseRepository struct {
}

func (p *ProjectsDatabaseRepository) GetAll() ([]entities.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectsDatabaseRepository) CreateOrUpdate(project entities.Project) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProjectsDatabaseRepository) Delete(uuid string) error {
	//TODO implement me
	panic("implement me")
}
