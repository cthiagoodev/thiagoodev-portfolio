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

	var about entities.About
	err := repository.pool.QueryRow(ctx, query).Scan(
		&about.UUID,
		&about.Name,
		&about.Phone,
		&about.Description,
		&about.Text,
		&about.CreatedAt,
		&about.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domainerrors.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get about: %w", err)
	}

	return &about, nil
}
