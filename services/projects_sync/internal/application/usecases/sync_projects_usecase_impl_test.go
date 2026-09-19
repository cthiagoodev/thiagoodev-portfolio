package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	githubmocks "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks/github"
	repositoriesmocks "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks/repositories"
	supabasemocks "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks/supabase"
	"github.com/stretchr/testify/require"
)

func TestSyncProjectsUseCaseImpl_Execute(t *testing.T) {
	t.Run("sync persisted projects to supabase", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		projects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(projects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, projects).Return(nil)
		replace.Once().NotBefore(fetch.Call)
		// Supabase must receive the persisted data, including database-generated fields.
		persistedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio",
			CreatedAt: time.Date(2026, time.September, 15, 12, 0, 0, 0, time.UTC)}}
		read := repository.EXPECT().GetAll(ctx).Return(persistedProjects, nil)
		read.Once().NotBefore(replace.Call)
		sync := supabaseService.EXPECT().ReplaceAll(ctx, persistedProjects).Return(nil)
		sync.Once().NotBefore(read.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.NoError(t, err)
	})

	t.Run("reject empty github results", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		githubService.EXPECT().FetchRepositories(ctx).Return([]entities.Project{}, nil).Once()
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.Error(t, err)
	})

	t.Run("reject nil github results", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		githubService.EXPECT().FetchRepositories(ctx).Return(nil, nil).Once()
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.Error(t, err)
	})

	t.Run("stop when github fails", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		expectedErr := errors.New("github failure")
		githubService.EXPECT().FetchRepositories(ctx).Return(nil, expectedErr).Once()
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("stop when replacement fails", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		expectedErr := errors.New("replace failure")
		projects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(projects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, projects).Return(expectedErr)
		replace.Once().NotBefore(fetch.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("stop when reading persisted projects fails", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		expectedErr := errors.New("read failure")
		projects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(projects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, projects).Return(nil)
		replace.Once().NotBefore(fetch.Call)
		read := repository.EXPECT().GetAll(ctx).Return(nil, expectedErr)
		read.Once().NotBefore(replace.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
	})

	t.Run("propagate supabase failure", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		expectedErr := errors.New("supabase failure")
		projects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(projects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, projects).Return(nil)
		replace.Once().NotBefore(fetch.Call)
		// Supabase must receive the persisted data, including database-generated fields.
		persistedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio",
			CreatedAt: time.Date(2026, time.September, 15, 12, 0, 0, 0, time.UTC)}}
		read := repository.EXPECT().GetAll(ctx).Return(persistedProjects, nil)
		read.Once().NotBefore(replace.Call)
		sync := supabaseService.EXPECT().ReplaceAll(ctx, persistedProjects).Return(expectedErr)
		sync.Once().NotBefore(read.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
	})

}
