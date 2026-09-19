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
}

func NewSyncProjectsUseCaseImpl(
	repository repositories.ProjectsRepository,
	githubService github.GithubService,
	supabaseService supabase.SupabaseService,
) *SyncProjectsUseCaseImpl {
	return &SyncProjectsUseCaseImpl{
		repository,
		githubService,
		supabaseService,
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

	cErr := s.repository.ResetAndCreateAll(ctx, projects)
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
