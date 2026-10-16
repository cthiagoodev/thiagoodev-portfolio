package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/handlers"
	mocks "github.com/cthiagoodev/portfolio/services/app-api/internal/test/mocks/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSocialContactHandler_GetAll_Success(t *testing.T) {
	ctx := context.Background()
	getAllSocialContactsUseCaseMock := mocks.NewMockGetAllSocialContactsUseCase(t)
	firstContact, err := entities.NewSocialContact(
		"social-uuid-1",
		"github",
		"https://github.com/cthiagoodev",
		"/images/icons/github.svg",
		time.Time{},
		time.Time{},
	)
	require.NoError(t, err)
	expectedContacts := []entities.SocialContact{*firstContact}
	getAllSocialContactsUseCaseMock.EXPECT().Execute(ctx).Return(expectedContacts, nil).Once()
	socialContactHandler := handlers.NewSocialContactHandler(getAllSocialContactsUseCaseMock)
	request := httptest.NewRequest(http.MethodGet, "/social-contacts", nil).WithContext(ctx)
	responseRecorder := httptest.NewRecorder()

	socialContactHandler.GetAll(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.Contains(t, responseRecorder.Body.String(), "github")
	assert.Contains(t, responseRecorder.Body.String(), "https://github.com/cthiagoodev")
}

func TestSocialContactHandler_GetAll_EmptyList(t *testing.T) {
	ctx := context.Background()
	getAllSocialContactsUseCaseMock := mocks.NewMockGetAllSocialContactsUseCase(t)
	expectedContacts := []entities.SocialContact{}
	getAllSocialContactsUseCaseMock.EXPECT().Execute(ctx).Return(expectedContacts, nil).Once()
	socialContactHandler := handlers.NewSocialContactHandler(getAllSocialContactsUseCaseMock)
	request := httptest.NewRequest(http.MethodGet, "/social-contacts", nil).WithContext(ctx)
	responseRecorder := httptest.NewRecorder()

	socialContactHandler.GetAll(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.Equal(t, "[]\n", responseRecorder.Body.String())
}

func TestSocialContactHandler_GetAll_InternalServerError(t *testing.T) {
	ctx := context.Background()
	databaseError := errors.New("database connection failed")
	getAllSocialContactsUseCaseMock := mocks.NewMockGetAllSocialContactsUseCase(t)
	getAllSocialContactsUseCaseMock.EXPECT().Execute(ctx).Return(nil, databaseError).Once()
	socialContactHandler := handlers.NewSocialContactHandler(getAllSocialContactsUseCaseMock)
	request := httptest.NewRequest(http.MethodGet, "/social-contacts", nil).WithContext(ctx)
	responseRecorder := httptest.NewRecorder()

	socialContactHandler.GetAll(responseRecorder, request)

	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.Contains(t, responseRecorder.Body.String(), "internal server error")
}

