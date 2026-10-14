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

func TestNewSkill_Success(t *testing.T) {
	label := "Golang"
	url := "https://go.dev"
	imagePath := "/images/skills/golang.svg"

	skill, err := entities.NewSkill(
		"",
		label,
		url,
		imagePath,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, label, skill.Label)
	assert.Equal(t, url, skill.URL)
	assert.Equal(t, imagePath, skill.ImagePath)
	assert.False(t, skill.CreatedAt.IsZero())
	assert.False(t, skill.UpdatedAt.IsZero())
}

func TestNewSkill_ValidationError_EmptyLabel(t *testing.T) {
	skill, err := entities.NewSkill(
		"",
		"   ",
		"https://go.dev",
		"/images/skills/golang.svg",
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, skill)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "label is required")
}
