package about

import (
	"context"
	"database/sql"

	"github.com/cthiagoodev/thiagoodev-portfolio/server/internal/common"
)

type DatabaseRepository struct {
	db *sql.DB
}

func NewDatabaseRepository(db *sql.DB) *DatabaseRepository {
	return &DatabaseRepository{db: db}
}

func (r *DatabaseRepository) Find(ctx context.Context) (About, error) {
	return About{}, common.ErrNotFound
}
