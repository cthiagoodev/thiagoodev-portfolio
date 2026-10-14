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

func TestNewEducation_Success_WithEndDate(t *testing.T) {
	startDate := time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC)

	education, err := entities.NewEducation(
		"",
		"Ciência da Computação",
		"Bacharelado em Ciência da Computação",
		"Universidade Federal",
		startDate,
		&endDate,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Equal(t, "Ciência da Computação", education.Course)
	assert.Equal(t, "Bacharelado em Ciência da Computação", education.Description)
	assert.Equal(t, "Universidade Federal", education.EducationalInstitution)
	assert.Equal(t, startDate, education.StartDate)
	assert.Equal(t, &endDate, education.EndDate)
	assert.False(t, education.CreatedAt.IsZero())
	assert.False(t, education.UpdatedAt.IsZero())
}

func TestNewEducation_Success_Ongoing_NilEndDate(t *testing.T) {
	startDate := time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC)

	education, err := entities.NewEducation(
		"",
		"Pós-Graduação em Arquitetura de Software",
		"Especialização em sistemas distribuídos",
		"Instituto de Tecnologia",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.NoError(t, err)
	assert.Nil(t, education.EndDate)
}

func TestNewEducation_ValidationError_EmptyCourse(t *testing.T) {
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	education, err := entities.NewEducation(
		"",
		"   ",
		"Descrição",
		"Instituição",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, education)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "course is required")
}

func TestNewEducation_ValidationError_EmptyDescription(t *testing.T) {
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	education, err := entities.NewEducation(
		"",
		"Curso",
		"   ",
		"Instituição",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, education)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "description is required")
}

func TestNewEducation_ValidationError_EmptyInstitution(t *testing.T) {
	startDate := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	education, err := entities.NewEducation(
		"",
		"Curso",
		"Descrição",
		"   ",
		startDate,
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, education)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "educational institution is required")
}

func TestNewEducation_ValidationError_ZeroStartDate(t *testing.T) {
	education, err := entities.NewEducation(
		"",
		"Curso",
		"Descrição",
		"Instituição",
		time.Time{},
		nil,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, education)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "start date is required")
}

func TestNewEducation_ValidationError_EndDateBeforeStartDate(t *testing.T) {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	education, err := entities.NewEducation(
		"",
		"Curso",
		"Descrição",
		"Instituição",
		startDate,
		&endDate,
		time.Time{},
		time.Time{},
	)

	require.Error(t, err)
	assert.Nil(t, education)
	assert.True(t, errors.Is(err, domainerrors.ErrDomainEntityValidation))
	assert.Contains(t, err.Error(), "end date cannot be before start date")
}
