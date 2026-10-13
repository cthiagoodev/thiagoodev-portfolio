package github

import (
	"testing"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapGithubProjectToEntity(t *testing.T) {
	t.Run("should map a github project to a domain entity", func(t *testing.T) {
		description := "Portfolio project"
		homepage := "https://github.com/cthiagoodev/portfolio"

		createdAt := time.Date(
			2026,
			time.September,
			6,
			12,
			0,
			0,
			0,
			time.UTC,
		)
		updatedAt := createdAt.Add(time.Hour)
		pushedAt := createdAt.Add(2 * time.Hour)

		input := Project{
			Id:          123,
			Name:        "portfolio",
			Description: &description,
			Homepage:    &homepage,
			Language:    "Go",
			CreatedAt:   createdAt,
			UpdatedAt:   &updatedAt,
			PushedAt:    &pushedAt,
		}

		result := MapGithubProjectToEntity(input)

		expected := entities.Project{
			Uuid:        "",
			ExternalId:  "123",
			Name:        "portfolio",
			Description: description,
			Url:         homepage,
			Languages:   []string{"Go"},
			CreatedAt:   createdAt,
			UpdatedAt:   &updatedAt,
			PushedAt:    &pushedAt,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("should map a github project when optional fields are nil", func(t *testing.T) {
		createdAt := time.Date(
			2026,
			time.September,
			6,
			12,
			0,
			0,
			0,
			time.UTC,
		)

		input := Project{
			Id:        123,
			Name:      "portfolio",
			Language:  "Go",
			CreatedAt: createdAt,
		}

		result := MapGithubProjectToEntity(input)

		expected := entities.Project{
			Uuid:        "",
			ExternalId:  "123",
			Name:        "portfolio",
			Description: "",
			Url:         "",
			Languages:   []string{"Go"},
			CreatedAt:   createdAt,
		}

		assert.Equal(t, expected, result)
	})
}

func TestMapGithubProjectsToEntities(t *testing.T) {
	t.Run("should return an empty list when input is empty", func(t *testing.T) {
		result := MapGithubProjectsToEntities([]Project{})

		require.NotNil(t, result)
		assert.Empty(t, result)
	})

	t.Run("should map multiple github projects", func(t *testing.T) {
		createdAt := time.Date(2026, time.September, 6, 12, 0, 0, 0, time.UTC)
		input := []Project{
			{Id: 1, Name: "repo1", Language: "Go", CreatedAt: createdAt},
			{Id: 2, Name: "repo2", Language: "TypeScript", CreatedAt: createdAt},
		}

		result := MapGithubProjectsToEntities(input)

		require.Len(t, result, 2)
		assert.Equal(t, "1", result[0].ExternalId)
		assert.Equal(t, "repo1", result[0].Name)
		assert.Equal(t, "2", result[1].ExternalId)
		assert.Equal(t, "repo2", result[1].Name)
	})
}
