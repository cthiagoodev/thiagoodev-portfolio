package github

import (
	"testing"
	"time"

	"github.com/cthiagoodev/thiagoodev-portfolio/services/projects_sync/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectsMapper(t *testing.T) {
	t.Run("should map github projects to domain projects", func(t *testing.T) {
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

		input := []Project{
			{
				Id:          123,
				Name:        "portfolio",
				Description: &description,
				HtmlUrl:     "https://github.com/cthiagoodev/portfolio",
				Languages:   []string{"Go", "SQL"},
				PushedAt:    pushedAt,
			},
		}

		before := time.Now()

		result := ProjectsMapper(input)

		after := time.Now()

		require.Len(t, result, 1)

		expected := entities.Project{
			Uuid:        "",
			ExternalId:  "123",
			Name:        "portfolio",
			Description: &description,
			Url:         "https://github.com/cthiagoodev/portfolio",
			Languages:   []string{"Go", "SQL"},
			UpdatedAt:   pushedAt,
		}

		assert.Equal(t, expected.Uuid, result[0].Uuid)
		assert.Equal(t, expected.ExternalId, result[0].ExternalId)
		assert.Equal(t, expected.Name, result[0].Name)
		assert.Equal(t, expected.Description, result[0].Description)
		assert.Equal(t, expected.Url, result[0].Url)
		assert.Equal(t, expected.Languages, result[0].Languages)
		assert.Equal(t, expected.UpdatedAt, result[0].UpdatedAt)

		assert.False(t, result[0].CreatedAt.Before(before))
		assert.False(t, result[0].CreatedAt.After(after))
	})

	t.Run("should return empty list when input is empty", func(t *testing.T) {
		result := ProjectsMapper([]Project{})

		require.NotNil(t, result)
		assert.Empty(t, result)
	})
}
