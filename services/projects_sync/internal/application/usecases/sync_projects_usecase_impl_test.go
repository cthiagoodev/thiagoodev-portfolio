package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
	githubmocks "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks/github"
	repositoriesmocks "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks/repositories"
	supabasemocks "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/test/mocks/supabase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSyncProjectsUseCaseImpl_Execute(t *testing.T) {
	t.Run("sync persisted projects to supabase", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		githubProjects := []github.Project{{Id: 123, Name: "portfolio"}}
		mappedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		mapperCalled := false
		mapper := func(repos []github.Project) []entities.Project {
			mapperCalled = true
			assert.Equal(t, githubProjects, repos)
			return mappedProjects
		}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(githubProjects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, mappedProjects).Return(nil)
		replace.Once().NotBefore(fetch.Call)
		// Supabase must receive the persisted data, including database-generated fields.
		persistedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio",
			CreatedAt: time.Date(2026, time.September, 15, 12, 0, 0, 0, time.UTC)}}
		read := repository.EXPECT().GetAll(ctx).Return(persistedProjects, nil)
		read.Once().NotBefore(replace.Call)
		sync := supabaseService.EXPECT().ReplaceAll(ctx, persistedProjects).Return(nil)
		sync.Once().NotBefore(read.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
		err := useCase.Execute(ctx)
		require.NoError(t, err)
		assert.True(t, mapperCalled)
	})

	t.Run("reject empty github results", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		mapper := func(repos []github.Project) []entities.Project {
			t.Fatal("mapper must not run after a failed or empty fetch")
			return nil
		}
		githubService.EXPECT().FetchRepositories(ctx).Return([]github.Project{}, nil).Once()
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
		err := useCase.Execute(ctx)
		require.Error(t, err)
	})

	t.Run("reject nil github results", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		mapper := func(repos []github.Project) []entities.Project {
			t.Fatal("mapper must not run after a failed or empty fetch")
			return nil
		}
		githubService.EXPECT().FetchRepositories(ctx).Return(nil, nil).Once()
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
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
		mapper := func(repos []github.Project) []entities.Project {
			t.Fatal("mapper must not run after a failed or empty fetch")
			return nil
		}
		githubService.EXPECT().FetchRepositories(ctx).Return(nil, expectedErr).Once()
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
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
		githubProjects := []github.Project{{Id: 123, Name: "portfolio"}}
		mappedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		mapperCalled := false
		mapper := func(repos []github.Project) []entities.Project {
			mapperCalled = true
			assert.Equal(t, githubProjects, repos)
			return mappedProjects
		}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(githubProjects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, mappedProjects).Return(expectedErr)
		replace.Once().NotBefore(fetch.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
		assert.True(t, mapperCalled)
	})

	t.Run("stop when reading persisted projects fails", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		expectedErr := errors.New("read failure")
		githubProjects := []github.Project{{Id: 123, Name: "portfolio"}}
		mappedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		mapperCalled := false
		mapper := func(repos []github.Project) []entities.Project {
			mapperCalled = true
			assert.Equal(t, githubProjects, repos)
			return mappedProjects
		}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(githubProjects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, mappedProjects).Return(nil)
		replace.Once().NotBefore(fetch.Call)
		read := repository.EXPECT().GetAll(ctx).Return(nil, expectedErr)
		read.Once().NotBefore(replace.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
		assert.True(t, mapperCalled)
	})

	t.Run("propagate supabase failure", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		repository := repositoriesmocks.NewMockProjectsRepository(t)
		githubService := githubmocks.NewMockGithubService(t)
		supabaseService := supabasemocks.NewMockSupabaseService(t)
		expectedErr := errors.New("supabase failure")
		githubProjects := []github.Project{{Id: 123, Name: "portfolio"}}
		mappedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio"}}
		mapperCalled := false
		mapper := func(repos []github.Project) []entities.Project {
			mapperCalled = true
			assert.Equal(t, githubProjects, repos)
			return mappedProjects
		}
		fetch := githubService.EXPECT().FetchRepositories(ctx).Return(githubProjects, nil)
		fetch.Once()
		replace := repository.EXPECT().ResetAndCreateAll(ctx, mappedProjects).Return(nil)
		replace.Once().NotBefore(fetch.Call)
		// Supabase must receive the persisted data, including database-generated fields.
		persistedProjects := []entities.Project{{ExternalId: "123", Name: "portfolio",
			CreatedAt: time.Date(2026, time.September, 15, 12, 0, 0, 0, time.UTC)}}
		read := repository.EXPECT().GetAll(ctx).Return(persistedProjects, nil)
		read.Once().NotBefore(replace.Call)
		sync := supabaseService.EXPECT().ReplaceAll(ctx, persistedProjects).Return(expectedErr)
		sync.Once().NotBefore(read.Call)
		useCase := NewSyncProjectsUseCaseImpl(repository, githubService, supabaseService, mapper)
		err := useCase.Execute(ctx)
		require.ErrorIs(t, err, expectedErr)
		assert.True(t, mapperCalled)
	})

}
