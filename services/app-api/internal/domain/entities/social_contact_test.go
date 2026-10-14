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

func TestNewSocialContact_Success(t *testing.T) {
	label := "github"
	url := "https://github.com/cthiagoodev"
	imagePath := "/images/icons/github.svg"

	contact, err := entities.NewSocialContact(
		"",
		label,
		url,
		imagePath,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, label, contact.Label)
	assert.Equal(t, url, contact.URL)
	assert.Equal(t, imagePath, contact.ImagePath)
	assert.False(t, contact.CreatedAt.IsZero())
	assert.False(t, contact.UpdatedAt.IsZero())
}

func TestNewSocialContact_ValidationError_EmptyLabel(t *testing.T) {
	contact, err := entities.NewSocialContact(
		"",
		"   ",
		"https://github.com/cthiagoodev",
		"",
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, contact)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "label is required")
}

func TestNewSocialContact_ValidationError_EmptyURL(t *testing.T) {
	contact, err := entities.NewSocialContact(
		"",
		"github",
		"   ",
		"",
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, contact)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "url is required")
}
