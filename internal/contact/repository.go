package contact

import "context"

type Repository interface {
	Find(ctx context.Context) (Contact, error)
}
