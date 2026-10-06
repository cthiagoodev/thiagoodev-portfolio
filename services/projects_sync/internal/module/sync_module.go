package module

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/application/usecases"
	domainusecases "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/usecases"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/database"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/repositories"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/supabase"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SyncModule struct {
	UseCase domainusecases.SyncProjectsUseCase
	DbPool  *pgxpool.Pool
	SbPool  *pgxpool.Pool
}

func NewSyncModule(ctx context.Context) (*SyncModule, error) {
	dbpool, err := database.NewPool(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	sbpool, err := supabase.NewPool(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create supabase pool: %w", err)
	}

	url, err := url.Parse(os.Getenv("GITHUB_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse GITHUB_URL: %w", err)
	}

	projectsDatabaseRepository := repositories.NewProjectsDatabaseRepository(dbpool)
	projectsGateway := supabase.NewGateway(sbpool)
	githubService := github.NewGithubServiceImpl(http.DefaultClient, url)

	syncProjectsUseCase := usecases.NewSyncProjectsUseCaseImpl(
		projectsDatabaseRepository,
		projectsGateway,
		githubService,
	)

	return &SyncModule{
		UseCase: syncProjectsUseCase,
		DbPool:  dbpool,
		SbPool:  sbpool,
	}, nil
}
