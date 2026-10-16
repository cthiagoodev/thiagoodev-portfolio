package repositories

import (
	"context"
	"fmt"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	domainrepositories "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/repositories"
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

	socialContacts := []entities.SocialContact{}
	for rows.Next() {
		var socialContact entities.SocialContact
		if err := rows.Scan(
			&socialContact.UUID,
			&socialContact.Label,
			&socialContact.URL,
			&socialContact.ImagePath,
			&socialContact.CreatedAt,
			&socialContact.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan social contact: %w", err)
		}
		socialContacts = append(socialContacts, socialContact)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during social contacts rows iteration: %w", err)
	}

	return socialContacts, nil
}

