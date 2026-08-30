package repositories

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

type SkillsRepository interface {
	GetByName(name string) (entities.Skill, error)
}
