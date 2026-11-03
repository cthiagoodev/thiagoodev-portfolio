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

var _ domainrepositories.SocialContactRepository = (*SocialContactDatabaseRepository)(nil)

type SocialContactDatabaseRepository struct {
	pool *pgxpool.Pool
}

func NewSocialContactDatabaseRepository(pool *pgxpool.Pool) *SocialContactDatabaseRepository {
	return &SocialContactDatabaseRepository{
		pool: pool,
	}
}

func (repository *SocialContactDatabaseRepository) GetAll(ctx context.Context) ([]entities.SocialContact, error) {
	query := `
		SELECT 
			uuid, 
			label, 
			url, 
			COALESCE(image_path, ''), 
			created_at, 
			updated_at
		FROM social_contacts
	`

	rows, err := repository.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query social contacts: %w", err)
	}
	defer rows.Close()

	socialContacts, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[entities.SocialContact],
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no sarching social contants %w", domainerrors.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to collect social contacts rows: %w", err)
	}

	if len(socialContacts) == 0 {
		return nil, fmt.Errorf("no sarching social contants %w", domainerrors.ErrNotFound)
	}

	return socialContacts, nil
}
