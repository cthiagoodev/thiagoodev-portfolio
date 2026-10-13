package entities_test

import (
	"errors"
	"testing"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProject_Success(t *testing.T) {
	name := "Portfolio Web"
	repoUrl := "https://github.com/cthiagoodev/portfolio"

	project, err := entities.NewProject(
		"",
		nil,
		name,
		"Meu portfólio pessoal",
		"https://thiagoodev.com",
		repoUrl,
		nil,
		nil,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, name, project.Name)
	assert.Equal(t, repoUrl, project.RepoUrl)
	assert.NotNil(t, project.Languages)
	assert.Empty(t, project.Languages)
	assert.False(t, project.CreatedAt.IsZero())
	assert.False(t, project.UpdatedAt.IsZero())
}

func TestNewProject_ValidationError_EmptyName(t *testing.T) {
	project, err := entities.NewProject(
		"",
		nil,
		"   ",
		"Descrição",
		"",
		"",
		[]string{"Go"},
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, project)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "project name is required")
}

