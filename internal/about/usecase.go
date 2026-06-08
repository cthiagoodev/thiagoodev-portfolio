package about

import "context"

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) Get(ctx context.Context) (About, error) {
	return u.repo.Find(ctx)
}
