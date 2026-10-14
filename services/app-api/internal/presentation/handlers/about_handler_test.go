package handlers_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/entities"
	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/handlers"
	mocks "github.com/cthiagoodev/portfolio/services/app-api/internal/test/mocks/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAboutHandler_Get_Success(t *testing.T) {
	ctx := context.Background()
	getAboutUseCaseMock := mocks.NewMockGetAboutUseCase(t)
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
	getAboutUseCaseMock.EXPECT().Execute(ctx).Return(expectedAbout, nil).Once()
	aboutHandler := handlers.NewAboutHandler(getAboutUseCaseMock)
	request := httptest.NewRequest(http.MethodGet, "/api/about", nil).WithContext(ctx)
	responseRecorder := httptest.NewRecorder()

	aboutHandler.Get(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.Contains(t, responseRecorder.Body.String(), "Thiago Sousa")
}

func TestAboutHandler_Get_NotFound(t *testing.T) {
	ctx := context.Background()
	getAboutUseCaseMock := mocks.NewMockGetAboutUseCase(t)
	getAboutUseCaseMock.EXPECT().Execute(ctx).Return(nil, domainerrors.ErrNotFound).Once()
	aboutHandler := handlers.NewAboutHandler(getAboutUseCaseMock)
	request := httptest.NewRequest(http.MethodGet, "/api/about", nil).WithContext(ctx)
	responseRecorder := httptest.NewRecorder()

	aboutHandler.Get(responseRecorder, request)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), "about not found")
}

func TestAboutHandler_Get_InternalServerError(t *testing.T) {
	ctx := context.Background()
	databaseError := errors.New("database connection failed")
	getAboutUseCaseMock := mocks.NewMockGetAboutUseCase(t)
	getAboutUseCaseMock.EXPECT().Execute(ctx).Return(nil, databaseError).Once()
	aboutHandler := handlers.NewAboutHandler(getAboutUseCaseMock)
	request := httptest.NewRequest(http.MethodGet, "/api/about", nil).WithContext(ctx)
	responseRecorder := httptest.NewRecorder()

	aboutHandler.Get(responseRecorder, request)

	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Contains(t, responseRecorder.Body.String(), "internal server error")
}
