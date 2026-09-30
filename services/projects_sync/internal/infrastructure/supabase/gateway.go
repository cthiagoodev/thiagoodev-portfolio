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

type Gateway struct {
	pool *pgxpool.Pool
}

func NewGateway(pool *pgxpool.Pool) *Gateway {
	return &Gateway{
		pool: pool,
	}
}

func (g *Gateway) PublishProjects(ctx context.Context, projects []entities.Project) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	err := pgx.BeginFunc(ctx, g.pool, func(tx pgx.Tx) error {
		return g.replaceProjects(ctx, tx, projects)
	})

	if err != nil {
		return fmt.Errorf("replace all Supabase projects: %w", err)
	}

	return nil
}

func (g *Gateway) replaceProjects(ctx context.Context, tx pgx.Tx, projects []entities.Project) error {
	if _, err := tx.Exec(ctx, "DELETE FROM projects"); err != nil {
		return fmt.Errorf("delete existing Supabase projects: %w", err)
	}

	if _, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"projects"},
		[]string{
			"uuid",
			"name",
			"description",
			"url",
			"start_date",
			"end_date",
			"created_at",
			"updated_at",
		},
		pgx.CopyFromSlice(len(projects), func(i int) ([]any, error) {
			return g.projectToRow(projects[i]), nil
		}),
	); err != nil {
		return fmt.Errorf("copy replacement Supabase projects: %w", err)
	}

	return g.replaceProjectSkills(ctx, tx, projects)
}

func (g *Gateway) replaceProjectSkills(ctx context.Context, tx pgx.Tx, projects []entities.Project) error {
	if _, err := tx.Exec(ctx, "DELETE FROM projects_skills"); err != nil {
		return fmt.Errorf("delete existing Supabase projects_skills: %w", err)
	}

	skills, err := g.loadSkills(ctx, tx)
	if err != nil {
		return err
	}

	projectSkills := g.collectProjectSkills(projects, skills)

	if _, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"projects_skills"},
		[]string{"project_id", "skill_id"},
		pgx.CopyFromSlice(len(projectSkills), func(i int) ([]any, error) {
			return g.projectSkillToRow(projectSkills[i]), nil
		}),
	); err != nil {
		return fmt.Errorf("copy replacement Supabase projects_skills: %w", err)
	}

	return nil
}

func (g *Gateway) loadSkills(ctx context.Context, tx pgx.Tx) (map[string]Skill, error) {
	rows, err := tx.Query(ctx, "SELECT uuid, label, url, image_path, created_at, updated_at FROM skills")
	if err != nil {
		return nil, fmt.Errorf("get existing Supabase skills: %w", err)
	}

	defer rows.Close()

	skills, err := pgx.CollectRows(
		rows,
		pgx.RowToStructByName[Skill],
	)
	if err != nil {
		return nil, fmt.Errorf("collect existing Supabase skills: %w", err)
	}

	dict := make(map[string]Skill, len(skills))

	for _, skill := range skills {
		dict[strings.TrimSpace(skill.Label)] = skill
	}

	return dict, nil
}

func (g *Gateway) projectToRow(project entities.Project) []any {
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

func (g *Gateway) collectProjectSkills(projects []entities.Project, skills map[string]Skill) []ProjectSkill {
	data := make([]ProjectSkill, 0)

	for _, project := range projects {
		for _, lang := range project.Languages {
			skill, ok := skills[strings.TrimSpace(lang)]

			if !ok {
				continue
			}

			data = append(data, ProjectSkill{
				ProjectId: project.Uuid,
				SkillId:   skill.Uuid,
			})
		}
	}

	return data
}

func (g *Gateway) projectSkillToRow(projectSkill ProjectSkill) []any {
	return []any{
		projectSkill.ProjectId,
		projectSkill.SkillId,
	}
}
