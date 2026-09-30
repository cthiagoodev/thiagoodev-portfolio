package usecases

import (
	"context"
	"fmt"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/gateways"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/repositories"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
)

type SyncProjectsUseCaseImpl struct {
	projectsRepository repositories.ProjectsRepository
	projectsGateway    gateways.ProjectsGateway
	githubService      github.GithubService
}

func NewSyncProjectsUseCaseImpl(
	projectsRepository repositories.ProjectsRepository,
	projectsGateway gateways.ProjectsGateway,
	githubService github.GithubService,
) *SyncProjectsUseCaseImpl {
	return &SyncProjectsUseCaseImpl{
		projectsRepository,
		projectsGateway,
		githubService,
	}
}

func (s *SyncProjectsUseCaseImpl) Execute(ctx context.Context) error {
	projects, gErr := s.githubService.FetchRepositories(ctx)
	if gErr != nil {
		return gErr
	}

	if len(projects) == 0 {
		return fmt.Errorf("there are no repositories to sync")
	}

	cErr := s.projectsRepository.ResetAndCreateAll(ctx, projects)
	if cErr != nil {
		return cErr
	}

	projects, pErr := s.projectsRepository.GetAll(ctx)
	if pErr != nil {
		return pErr
	}

	sErr := s.projectsGateway.PublishProjects(ctx, projects)
	return sErr
}
