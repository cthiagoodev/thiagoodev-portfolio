package usecases

import (
	"context"
	"fmt"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/repositories"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/supabase"
)

type SyncProjectsUseCaseImpl struct {
	repository      repositories.ProjectsRepository
	githubService   github.GithubService
	supabaseService supabase.SupabaseService
	mapper          github.ProjectsMapperFunc
}

func NewSyncProjectsUseCaseImpl(
	repository repositories.ProjectsRepository,
	githubService github.GithubService,
	supabaseService supabase.SupabaseService,
	mapper github.ProjectsMapperFunc,
) *SyncProjectsUseCaseImpl {
	return &SyncProjectsUseCaseImpl{
		repository,
		githubService,
		supabaseService,
		mapper,
	}
}

func (s *SyncProjectsUseCaseImpl) Execute(ctx context.Context) error {
	repos, gErr := s.githubService.FetchRepositories(ctx)
	if gErr != nil {
		return gErr
	}

	if len(repos) == 0 {
		return fmt.Errorf("there are no repositories to sync")
	}

	newProjects := s.mapper(repos)

	cErr := s.repository.ResetAndCreateAll(ctx, newProjects)
	if cErr != nil {
		return cErr
	}

	projects, pErr := s.repository.GetAll(ctx)
	if pErr != nil {
		return pErr
	}

	sErr := s.supabaseService.ReplaceAll(ctx, projects)
	return sErr
}
