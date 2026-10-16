package usecases_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/application/usecases"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	mocks "github.com/cthiagoodev/portfolio/services/app-api/internal/test/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllSocialContactsUseCase_Success(t *testing.T) {
	ctx := context.Background()
	socialContactRepositoryMock := mocks.NewMockSocialContactRepository(t)
	firstContact, err := entities.NewSocialContact(
		"social-uuid-1",
		"github",
		"https://github.com/cthiagoodev",
		"/images/icons/github.svg",
		time.Time{},
		time.Time{},
	)
	require.NoError(t, err)
	secondContact, err := entities.NewSocialContact(
		"social-uuid-2",
		"linkedin",
		"https://linkedin.com/in/cthiagoodev",
		"/images/icons/linkedin.svg",
		time.Time{},
		time.Time{},
	)
	require.NoError(t, err)
	expectedContacts := []entities.SocialContact{*firstContact, *secondContact}
	socialContactRepositoryMock.EXPECT().GetAll(ctx).Return(expectedContacts, nil).Once()
	useCase := usecases.NewGetAllSocialContactsUseCase(socialContactRepositoryMock)

	result, err := useCase.Execute(ctx)

	require.NoError(t, err)
	assert.Equal(t, expectedContacts, result)
}

func TestGetAllSocialContactsUseCase_EmptyList(t *testing.T) {
	ctx := context.Background()
	socialContactRepositoryMock := mocks.NewMockSocialContactRepository(t)
	expectedContacts := []entities.SocialContact{}
	socialContactRepositoryMock.EXPECT().GetAll(ctx).Return(expectedContacts, nil).Once()
	useCase := usecases.NewGetAllSocialContactsUseCase(socialContactRepositoryMock)

	result, err := useCase.Execute(ctx)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result)
}

func TestGetAllSocialContactsUseCase_RepositoryError(t *testing.T) {
	ctx := context.Background()
	databaseError := errors.New("database connection failed")
	socialContactRepositoryMock := mocks.NewMockSocialContactRepository(t)
	socialContactRepositoryMock.EXPECT().GetAll(ctx).Return(nil, databaseError).Once()
	useCase := usecases.NewGetAllSocialContactsUseCase(socialContactRepositoryMock)

	result, err := useCase.Execute(ctx)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, databaseError))
}

