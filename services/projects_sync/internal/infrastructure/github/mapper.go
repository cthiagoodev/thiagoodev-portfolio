package github

import (
	"strconv"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
)

func MapGithubProjectToEntity(project Project) entities.Project {
	return entities.Project{
		Uuid:        "",
		ExternalId:  strconv.FormatInt(project.Id, 10),
		Name:        project.Name,
		Description: project.Description,
		Url:         project.HtmlUrl,
		Languages:   project.Languages,
		CreatedAt:   time.Now(),
		UpdatedAt:   project.PushedAt,
	}
}

func MapGithubProjectsToEntities(projects []Project) []entities.Project {
	domainProjects := make([]entities.Project, len(projects))

	for i, project := range projects {
		domainProjects[i] = MapGithubProjectToEntity(project)
	}

	return domainProjects
}
