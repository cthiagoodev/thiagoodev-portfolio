package about

import "context"

type Repository interface {
	Find(ctx context.Context) (About, error)
}
