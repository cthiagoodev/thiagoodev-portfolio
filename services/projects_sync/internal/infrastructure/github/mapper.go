package github

import (
	"strconv"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/pkg/utils"
)

func MapGithubProjectToEntity(project Project) entities.Project {
	return entities.Project{
		Uuid:        "",
		ExternalId:  strconv.FormatInt(project.Id, 10),
		Name:        project.Name,
		Description: utils.DerefString(project.Description),
		Url:         utils.DerefString(project.Homepage),
		Languages:   []string{project.Language},
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		PushedAt:    project.PushedAt,
	}
}

func MapGithubProjectsToEntities(projects []Project) []entities.Project {
	domainProjects := make([]entities.Project, len(projects))

	for i, project := range projects {
		domainProjects[i] = MapGithubProjectToEntity(project)
	}

	return domainProjects
}
