package common

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func ParseDbError(err error) error {
	var connErr *pgconn.ConnectError
	if errors.As(err, &connErr) {
		return ErrConnection
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	return err
}
