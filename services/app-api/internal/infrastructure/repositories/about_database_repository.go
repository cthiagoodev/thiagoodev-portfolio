package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	domainrepositories "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/repositories"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ domainrepositories.AboutRepository = (*AboutDatabaseRepository)(nil)

type AboutDatabaseRepository struct {
	pool *pgxpool.Pool
}

func NewAboutDatabaseRepository(pool *pgxpool.Pool) *AboutDatabaseRepository {
	return &AboutDatabaseRepository{
		pool: pool,
	}
}

func (repository *AboutDatabaseRepository) Get(ctx context.Context) (*entities.About, error) {
	query := `
		SELECT 
			uuid, 
			name, 
			COALESCE(phone, ''), 
			COALESCE(description, ''), 
			COALESCE(text, ''), 
			created_at, 
			updated_at
		FROM about
		LIMIT 1
	`

	rows, err := repository.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get about: %w", err)
	}
	defer rows.Close()

	about, err := pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[entities.About],
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no sarching about %w", domainerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to collect about row: %w", err)
	}

	return &about, nil
}
