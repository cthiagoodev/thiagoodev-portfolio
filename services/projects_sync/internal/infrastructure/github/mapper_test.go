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

		pushedAt := time.Date(
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
			Id:          123,
			Name:        "portfolio",
			Description: &description,
			HtmlUrl:     "https://github.com/cthiagoodev/portfolio",
			Languages:   []string{"Go", "SQL"},
			PushedAt:    pushedAt,
		}

		before := time.Now()

		result := MapGithubProjectToEntity(input)

		after := time.Now()

		expected := entities.Project{
			Uuid:        "",
			ExternalId:  "123",
			Name:        "portfolio",
			Description: &description,
			Url:         "https://github.com/cthiagoodev/portfolio",
			Languages:   []string{"Go", "SQL"},
			UpdatedAt:   pushedAt,
		}

		assert.Equal(t, expected.Uuid, result.Uuid)
		assert.Equal(t, expected.ExternalId, result.ExternalId)
		assert.Equal(t, expected.Name, result.Name)
		assert.Equal(t, expected.Description, result.Description)
		assert.Equal(t, expected.Url, result.Url)
		assert.Equal(t, expected.Languages, result.Languages)
		assert.Equal(t, expected.UpdatedAt, result.UpdatedAt)

		assert.False(t, result.CreatedAt.Before(before))
		assert.False(t, result.CreatedAt.After(after))
	})
}

func TestMapGithubProjectsToEntities(t *testing.T) {
	t.Run("should return an empty list when input is empty", func(t *testing.T) {
		result := MapGithubProjectsToEntities([]Project{})

		require.NotNil(t, result)
		assert.Empty(t, result)
	})
}
