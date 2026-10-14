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
	repoURL := "https://github.com/cthiagoodev/portfolio"

	project, err := entities.NewProject(
		"",
		nil,
		name,
		"Meu portfólio pessoal",
		"https://thiagoodev.com",
		repoURL,
		nil,
		nil,
		nil,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, name, project.Name)
	assert.Equal(t, repoURL, project.RepoURL)
	assert.NotNil(t, project.Languages)
	assert.Empty(t, project.Languages)
	assert.NotNil(t, project.Skills)
	assert.Empty(t, project.Skills)
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
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, project)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "project name is required")
}

func TestProject_AddSkill_Success(t *testing.T) {
	project, err := entities.NewProject("", nil, "Projeto", "", "", "", nil, nil, nil, time.Time{}, time.Time{})
	require.NoError(t, err)
	skill, err := entities.NewSkill("skill-uuid-1", "Go", "https://go.dev", "", time.Time{}, time.Time{})
	require.NoError(t, err)

	project.AddSkill(*skill)

	assert.Len(t, project.Skills, 1)
	assert.Equal(t, "skill-uuid-1", project.Skills[0].UUID)
	assert.Equal(t, "Go", project.Skills[0].Label)
}

func TestProject_AddSkill_Idempotent(t *testing.T) {
	project, err := entities.NewProject("", nil, "Projeto", "", "", "", nil, nil, nil, time.Time{}, time.Time{})
	require.NoError(t, err)
	skill, err := entities.NewSkill("skill-uuid-1", "Go", "https://go.dev", "", time.Time{}, time.Time{})
	require.NoError(t, err)

	project.AddSkill(*skill)
	project.AddSkill(*skill)

	assert.Len(t, project.Skills, 1)
}

func TestProject_RemoveSkill_Success(t *testing.T) {
	skill1, err := entities.NewSkill("skill-uuid-1", "Go", "https://go.dev", "", time.Time{}, time.Time{})
	require.NoError(t, err)
	skill2, err := entities.NewSkill("skill-uuid-2", "Docker", "https://docker.com", "", time.Time{}, time.Time{})
	require.NoError(t, err)
	project, err := entities.NewProject("", nil, "Projeto", "", "", "", nil, []entities.Skill{*skill1, *skill2}, nil, time.Time{}, time.Time{})
	require.NoError(t, err)

	project.RemoveSkill("skill-uuid-1")

	assert.Len(t, project.Skills, 1)
	assert.Equal(t, "skill-uuid-2", project.Skills[0].UUID)
}
