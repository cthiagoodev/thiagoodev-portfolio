package supabase

import (
	"context"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
)

type SupabaseService interface {
	ReplaceAll(ctx context.Context, projects []entities.Project) error
}
