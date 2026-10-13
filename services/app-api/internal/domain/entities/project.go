package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type Project struct {
	Uuid        string
	ExternalId  *int64
	Name        string
	Description string
	Url         string
	RepoUrl     string
	Languages   []string
	PushedAt    *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(
	uuid string,
	externalId *int64,
	name string,
	description string,
	url string,
	repoUrl string,
	languages []string,
	pushedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Project, error) {
	if languages == nil {
		languages = []string{}
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	project := &Project{
		Uuid:        strings.TrimSpace(uuid),
		ExternalId:  externalId,
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Url:         strings.TrimSpace(url),
		RepoUrl:     strings.TrimSpace(repoUrl),
		Languages:   languages,
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
