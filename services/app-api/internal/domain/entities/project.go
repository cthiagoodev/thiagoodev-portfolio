package entities

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type Project struct {
	UUID        string
	ExternalID  *int64
	Name        string
	Description string
	URL         string
	RepoURL     string
	Languages   []string
	Skills      []Skill
	PushedAt    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(
	uuid string,
	externalID *int64,
	name string,
	description string,
	url string,
	repoURL string,
	languages []string,
	skills []Skill,
	pushedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Project, error) {
	if languages == nil {
		languages = []string{}
	}

	if skills == nil {
		skills = []Skill{}
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	project := &Project{
		UUID:        strings.TrimSpace(uuid),
		ExternalID:  externalID,
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		URL:         strings.TrimSpace(url),
		RepoURL:     strings.TrimSpace(repoURL),
		Languages:   languages,
		Skills:      skills,
		PushedAt:    pushedAt,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if err := project.Validate(); err != nil {
		return nil, err
	}

	return project, nil
}

func (p *Project) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("%w: project name is required", errors.ErrDomainEntityValidation)
	}
	return nil
}

func (p *Project) AddSkill(skill Skill) {
	exists := slices.ContainsFunc(p.Skills, func(s Skill) bool {
		return (skill.UUID != "" && s.UUID == skill.UUID) ||
			(skill.Label != "" && strings.EqualFold(s.Label, skill.Label))
	})
	if exists {
		return
	}
	p.Skills = append(p.Skills, skill)
}

func (p *Project) RemoveSkill(skillUUID string) {
	index := slices.IndexFunc(p.Skills, func(s Skill) bool {
		return s.UUID == skillUUID
	})
	if index != -1 {
		p.Skills = slices.Delete(p.Skills, index, index+1)
	}
}
