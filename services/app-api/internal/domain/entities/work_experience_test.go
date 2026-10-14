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

func TestNewWorkExperience_Success_WithEndDate(t *testing.T) {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	exp, err := entities.NewWorkExperience(
		"",
		"Software Engineer",
		"Desenvolvimento backend em Go",
		"Tech Company",
		startDate,
		&endDate,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, "Software Engineer", exp.Role)
	assert.Equal(t, "Desenvolvimento backend em Go", exp.Description)
	assert.Equal(t, "Tech Company", exp.Company)
	assert.Equal(t, startDate, exp.StartDate)
	assert.Equal(t, &endDate, exp.EndDate)
	assert.False(t, exp.CreatedAt.IsZero())
	assert.False(t, exp.UpdatedAt.IsZero())
}

func TestNewWorkExperience_Success_CurrentJob_NilEndDate(t *testing.T) {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	exp, err := entities.NewWorkExperience(
		"",
		"Senior Software Engineer",
		"Liderança técnica",
		"Global Tech",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Nil(t, exp.EndDate)
}

func TestNewWorkExperience_ValidationError_EmptyRole(t *testing.T) {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	exp, err := entities.NewWorkExperience(
		"",
		"   ",
		"Descrição",
		"Empresa",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, exp)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "role is required")
}

func TestNewWorkExperience_ValidationError_EmptyDescription(t *testing.T) {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	exp, err := entities.NewWorkExperience(
		"",
		"Cargo",
		"   ",
		"Empresa",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, exp)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "description is required")
}

func TestNewWorkExperience_ValidationError_EmptyCompany(t *testing.T) {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	exp, err := entities.NewWorkExperience(
		"",
		"Cargo",
		"Descrição",
		"   ",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, exp)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "company is required")
}

func TestNewWorkExperience_ValidationError_ZeroStartDate(t *testing.T) {
	exp, err := entities.NewWorkExperience(
		"",
		"Cargo",
		"Descrição",
		"Empresa",
		time.Time{},
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, exp)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "start date is required")
}

func TestNewWorkExperience_ValidationError_EndDateBeforeStartDate(t *testing.T) {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	exp, err := entities.NewWorkExperience(
		"",
		"Cargo",
		"Descrição",
		"Empresa",
		startDate,
		&endDate,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, exp)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "end date cannot be before start date")
}
