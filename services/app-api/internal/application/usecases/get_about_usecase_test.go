package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/application/usecases"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	mocks "github.com/cthiagoodev/portfolio/services/app-api/internal/test/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAboutUseCase_Success(t *testing.T) {
	ctx := context.Background()
	aboutRepositoryMock := mocks.NewMockAboutRepository(t)
	expectedAbout, err := entities.NewAbout(
		"about-uuid-1",
		"Thiago Sousa",
		"+5511999999999",
		"Software Engineer",
		"Bio completa",
		time.Time{},
		time.Time{},
	)
	require.NoError(t, err)
	aboutRepositoryMock.EXPECT().Get(ctx).Return(expectedAbout, nil).Once()
	useCase := usecases.NewGetAboutUseCase(aboutRepositoryMock)

	result, err := useCase.Execute(ctx)

	require.NoError(t, err)
	assert.Equal(t, expectedAbout, result)
}

func TestGetAboutUseCase_NotFound(t *testing.T) {
	ctx := context.Background()
	aboutRepositoryMock := mocks.NewMockAboutRepository(t)
	aboutRepositoryMock.EXPECT().Get(ctx).Return(nil, domainerrors.ErrNotFound).Once()
	useCase := usecases.NewGetAboutUseCase(aboutRepositoryMock)

	result, err := useCase.Execute(ctx)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, domainerrors.ErrNotFound))
}

func TestGetAboutUseCase_RepositoryError(t *testing.T) {
	ctx := context.Background()
	databaseError := errors.New("database connection failed")
	aboutRepositoryMock := mocks.NewMockAboutRepository(t)
	aboutRepositoryMock.EXPECT().Get(ctx).Return(nil, databaseError).Once()
	useCase := usecases.NewGetAboutUseCase(aboutRepositoryMock)

	result, err := useCase.Execute(ctx)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, databaseError))
}
