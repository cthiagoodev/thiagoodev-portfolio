package about

import (
	"context"
	"errors"
	"fmt"

	"github.com/cthiagoodev/thiagoodev-portfolio/server/internal/common"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseRepository struct {
	pool *pgxpool.Pool
}

func NewDatabaseRepository(pool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{pool}
}

func (r *DatabaseRepository) Find(ctx context.Context) (About, error) {
	query := "SELECT uuid, name, bio, photo, curriculum, linkedin, github, city, state FROM about LIMIT 1"
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
	)

	connectionErrorType := *pgconn.ConnectError
	if errors.As(err, &connectionErrorType) {
		return About{}, common.ErrConnection
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return About{}, common.ErrNotFound
	}

	if err != nil {
		return About{}, fmt.Errorf("about not found %w", err)
	}

	return about, nil
}
