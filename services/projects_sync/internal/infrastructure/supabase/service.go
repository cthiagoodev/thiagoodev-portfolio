package supabase

import (
	"context"
	"fmt"
	"strings"
	"time"

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
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)

	defer cancel()

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

		defer skillRows.Close()

		skills, err := pgx.CollectRows(
			skillRows,
			pgx.RowToStructByName[Skill],
		)

		if err != nil {
			return fmt.Errorf("collect existing Supabase skills: %w", err)
		}

		projectSkills := s.collectProjectSkills(projects, skills)

		if _, err := tx.CopyFrom(
			ctx,
			pgx.Identifier{"projects_skills"},
			[]string{"project_id", "skill_id"},
			pgx.CopyFromSlice(len(projectSkills), func(i int) ([]any, error) {
				return s.projectSkillToRow(projectSkills[i]), nil
			}),
		); err != nil {
			return fmt.Errorf("copy replacement Supabase projects_skills: %w", err)
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

func (s *SupabaseServiceImpl) collectProjectSkills(projects []entities.Project, skills []Skill) []ProjectSkill {
	data := make([]ProjectSkill, 0)

	for _, project := range projects {
		for _, lang := range project.Languages {
			for _, skill := range skills {
				if strings.TrimSpace(skill.Label) == strings.TrimSpace(lang) {
					data = append(data, ProjectSkill{
						ProjectId: project.Uuid,
						SkillId:   skill.Uuid,
					})
				}
			}
		}
	}

	return data
}

func (s *SupabaseServiceImpl) projectSkillToRow(projectSkill ProjectSkill) []any {
	return []any{
		projectSkill.ProjectId,
		projectSkill.SkillId,
	}
}
