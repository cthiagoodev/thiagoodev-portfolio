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

func TestNewAbout_Success(t *testing.T) {
	name := "Thiago Sousa"
	phone := "+5511999999999"
	description := "Software Engineer"
	text := "Bio completa sobre mim"

	about, err := entities.NewAbout(
		"",
		name,
		phone,
		description,
		text,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, name, about.Name)
	assert.Equal(t, phone, about.Phone)
	assert.Equal(t, description, about.Description)
	assert.Equal(t, text, about.Text)
	assert.False(t, about.CreatedAt.IsZero())
	assert.False(t, about.UpdatedAt.IsZero())
}

func TestNewAbout_ValidationError_EmptyName(t *testing.T) {
	about, err := entities.NewAbout(
		"",
		"   ",
		"+5511999999999",
		"Descrição",
		"Texto",
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, about)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "name is required")
}

func TestNewAbout_ValidationError_PhoneTooLong(t *testing.T) {
	about, err := entities.NewAbout(
		"",
		"Thiago",
		"+5511999999999999999999999",
		"Descrição",
		"Texto",
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, about)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "phone must have at most 20 characters")
}
