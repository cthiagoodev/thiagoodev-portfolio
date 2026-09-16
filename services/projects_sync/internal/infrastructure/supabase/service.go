package supabase

import (
	"context"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	sb "github.com/supabase-community/supabase-go"
)

type SupabaseService interface {
	ReplaceAll(ctx context.Context, projects []entities.Project) error
}

type SupabaseServiceImpl struct {
	client *sb.Client
}

func NewSupabaseService(client *sb.Client) SupabaseService {
	return &SupabaseServiceImpl{
		client: client,
	}
}

func (s *SupabaseServiceImpl) ReplaceAll(ctx context.Context, projects []entities.Project) error {
	return nil
}
