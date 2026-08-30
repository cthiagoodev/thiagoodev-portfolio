package usecases

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/repositories"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
)

type SyncProjectsUseCaseImpl struct {
	repository    repositories.ProjectsRepository
	githubService github.Service
}

func NewSyncProjectsUseCaseImpl(
	repository repositories.ProjectsRepository,
	githubService github.Service,
) *SyncProjectsUseCaseImpl {
	return &SyncProjectsUseCaseImpl{
		repository,
		githubService,
	}
}

func (s *SyncProjectsUseCaseImpl) Execute() ([]entities.Project, error) {
	repos, gErr := s.githubService.FetchRepositories()

	if gErr != nil {
		return nil, gErr
	}

	if len(repos) == 0 {
		return []entities.Project{}, nil
	}

	newProjects := github.ProjectsMapper(repos)

	dErr := s.repository.DeleteAll()

	if dErr != nil {
		return nil, dErr
	}

	projects, pErr := s.repository.CreateAll(newProjects)

	if pErr != nil {
		return nil, pErr
	}

	return projects, nil
}
