package supabase

import "github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"

func MapProjectToSupabaseModel(project entities.Project) Project {
	return Project{
		Uuid:        "",
		Name:        project.Name,
		Description: project.Description,
		Url:         &project.Url,
		StartDate:   project.CreatedAt,
		EndDate:     nil,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func MapProjectsToSupabaseModels(projects []entities.Project) []Project {
	supabaseProjects := make([]Project, len(projects))

	for i, project := range projects {
		supabaseProjects[i] = MapProjectToSupabaseModel(project)
	}

	return supabaseProjects
}
