package usecases

import (
	"context"
)

type SyncProjectsUseCase interface {
	Execute(ctx context.Context) error
}
