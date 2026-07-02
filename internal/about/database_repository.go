package about

import (
	"context"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseRepository struct {
	pool *pgxpool.Pool
}

func NewDatabaseRepository(pool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{pool}
}

func (r *DatabaseRepository) Find(ctx context.Context) (About, error) {
	about, aboutErr := r.findAbout(ctx)

	if aboutErr != nil {
		return About{}, aboutErr
	}

	techs, techsErr := r.findTechnologies(ctx, about)

	if techsErr != nil {
		return About{}, techsErr
	}

	about.AddStack(techs)

	return about, nil
}

func (r *DatabaseRepository) findAbout(ctx context.Context) (About, error) {
	query := "SELECT uuid, name, bio, photo, curriculum, linkedin, github, city, state, created_at, updated_at FROM about LIMIT 1"
	row := r.pool.QueryRow(ctx, query)

	about := About{}
	err := row.Scan(
		&about.Uuid,
		&about.Name,
		&about.Bio,
		&about.Photo,
		&about.Curriculum,
		&about.Linkedin,
		&about.GitHub,
		&about.City,
		&about.State,
		&about.CreatedAt,
		&about.UpdatedAt,
	)

	if err != nil {
		return About{}, common.ParseDbError(err)
	}

	return about, nil
}

func (r *DatabaseRepository) findTechnologies(ctx context.Context, about About) ([]Technology, error) {
	query := "SELECT uuid, name FROM technology WHERE about_uuid = $1"
	rows, err := r.pool.Query(ctx, query, about.Uuid)

	if err != nil {
		return []Technology{}, common.ParseDbError(err)
	}

	defer rows.Close()

	var techs []Technology

	for rows.Next() {
		tech := Technology{}
		sErr := rows.Scan(
			&tech.Uuid,
			&tech.Name,
		)

		if sErr != nil {
			return []Technology{}, common.ParseDbError(err)
		}

		techs = append(techs, tech)
	}

	if rErr := rows.Err(); rErr != nil {
		return nil, common.ParseDbError(err)
	}

	return techs, nil
}
