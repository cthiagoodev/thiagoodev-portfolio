package github

import (
	"strconv"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
)

func ProjectsMapper(ps []Project) []entities.Project {
	if len(ps) == 0 {
		return []entities.Project{}
	}

	projects := make([]entities.Project, len(ps))

	for _, p := range ps {
		projects = append(projects, entities.Project{
			Uuid:        "",
			ExternalId:  strconv.FormatInt(p.Id, 10),
			Name:        p.Name,
			Description: p.Description,
			Url:         p.HtmlUrl,
			Languages:   p.Languages,
			CreatedAt:   time.Time{},
			UpdatedAt:   p.PushedAt,
		})
	}

	return projects
}
