package usecases

import (
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/repositories"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
)

type SyncProjectsUseCaseImpl struct {
	repository    repositories.ProjectsRepository
	githubService github.GithubService
	mapper        github.ProjectsMapperFunc
}

func NewSyncProjectsUseCaseImpl(
	repository repositories.ProjectsRepository,
	githubService github.GithubService,
	mapper github.ProjectsMapperFunc,
) *SyncProjectsUseCaseImpl {
	return &SyncProjectsUseCaseImpl{
		repository,
		githubService,
		mapper,
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

	newProjects := s.mapper(repos)

	dErr := s.repository.DeleteAll()

	if dErr != nil {
		return nil, dErr
	}

	cErr := s.repository.CreateAll(newProjects)

	if cErr != nil {
		return nil, cErr
	}

	projects, pErr := s.repository.GetAll()

	if pErr != nil {
		return nil, pErr
	}

	return projects, nil
}
