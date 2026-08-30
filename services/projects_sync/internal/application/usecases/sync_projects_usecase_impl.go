package usecases

import (
	"strconv"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/repositories"
	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/infrastructure/github"
)

type SyncProjectsUseCaseImpl struct {
	projectsRepository repositories.ProjectsRepository
	skillsRepository   repositories.SkillsRepository
	githubService      github.Service
}

func NewSyncProjectsUseCaseImpl(
	projectsRepository repositories.ProjectsRepository,
	skillsRepository repositories.SkillsRepository,
	githubService github.Service,
) *SyncProjectsUseCaseImpl {
	return &SyncProjectsUseCaseImpl{
		projectsRepository,
		skillsRepository,
		githubService,
	}
}

func (s *SyncProjectsUseCaseImpl) Execute() ([]entities.Project, error) {
	repos, gErr := s.githubService.FetchRepositories()

	if gErr != nil {
		return nil, gErr
	}

	if len(repos) == 0 {
		return []entities.Project{}, nil
	}

	newProjects := make([]entities.Project, len(repos))

	for _, r := range repos {
		var skills []entities.Skill

		for _, l := range r.Languages {
			skill, sErr := s.skillsRepository.GetByName(l)

			if sErr != nil {
				continue
			}

			skills = append(skills, skill)
		}

		newProjects = append(newProjects, entities.Project{
			Uuid:        "",
			ExternalId:  strconv.FormatInt(r.Id, 10),
			Name:        r.Name,
			Description: r.Description,
			Url:         r.HtmlUrl,
			Skills:      skills,
			CreatedAt:   time.Time{},
			UpdatedAt:   r.PushedAt,
		})
	}

	projects, pErr := s.projectsRepository.CreateOrUpdateAll(newProjects)

	if pErr != nil {
		return nil, pErr
	}

	return projects, nil
}
