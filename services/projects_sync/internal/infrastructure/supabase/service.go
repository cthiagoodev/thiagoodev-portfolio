package supabase

import (
	"context"
	"fmt"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var projectColumns = []string{
	"uuid",
	"name",
	"description",
	"url",
	"start_date",
	"end_date",
	"created_at",
	"updated_at",
}

type SupabaseService interface {
	ReplaceAll(ctx context.Context, projects []entities.Project) error
}

type SupabaseServiceImpl struct {
	pool *pgxpool.Pool
}

func NewSupabaseService(pool *pgxpool.Pool) SupabaseService {
	return &SupabaseServiceImpl{
		pool: pool,
	}
}

func (s *SupabaseServiceImpl) ReplaceAll(ctx context.Context, projects []entities.Project) error {
	err := pgx.BeginFunc(ctx, s.pool, func(tx pgx.Tx) error {
		if _, err := tx.Exec(ctx, "DELETE FROM projects"); err != nil {
			return fmt.Errorf("delete existing Supabase projects: %w", err)
		}

		if _, err := tx.CopyFrom(
			ctx,
			pgx.Identifier{"projects"},
			projectColumns,
			pgx.CopyFromSlice(len(projects), func(i int) ([]any, error) {
				return s.projectToRow(projects[i]), nil
			}),
		); err != nil {
			return fmt.Errorf("copy replacement Supabase projects: %w", err)
		}

		if _, err := tx.Exec(ctx, "DELETE FROM projects_skills"); err != nil {
			return fmt.Errorf("delete existing Supabase projects_skills: %w", err)
		}

		skillRows, err := tx.Query(ctx, "SELECT uuid, label, url, image_path, created_at, updated_at FROM skills")
		if err != nil {
			return fmt.Errorf("get existing Supabase skills: %w", err)
		}

		skills, err := pgx.CollectRows(
			skillRows,
			pgx.RowToStructByName[Skill],
		)

		for _, project := range projects {
			for _, lang := range project.Languages {
				for _, skill := range skills {
					if skill.Label == lang {
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("replace all Supabase projects: %w", err)
	}

	return nil
}

func (s *SupabaseServiceImpl) projectToRow(project entities.Project) []any {
	return []any{
		project.Uuid,
		project.Name,
		project.Description,
		project.Url,
		project.CreatedAt,
		nil,
		project.CreatedAt,
		project.UpdatedAt,
	}
}

func (s *SupabaseServiceImpl) projectSkillToRow(projectUuid string, skillUuid string) []any {
	return []any{
		projectUuid,
		skillUuid,
	}
}
