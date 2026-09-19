package supabase

import (
	"context"
	"fmt"

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
	table := s.client.From("projects")

	_, _, err := table.Delete("", "").Execute()
	if err != nil {
		return fmt.Errorf("Error on delete all items in supabase projects table")
	}

	_, _, iErr := table.Insert(
		projects,
		true,
		"uuid",
		"",
		"",
	).Execute()
	if iErr != nil {
		return fmt.Errorf("Error on insert all items in supabase projects table")
	}

	//TODO: Implement the same for skills and project_skills tables

	return nil
}
