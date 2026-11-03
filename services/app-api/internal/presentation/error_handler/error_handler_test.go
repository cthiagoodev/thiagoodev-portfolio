package errorhandler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	errorhandler "github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/error_handler"
	"github.com/stretchr/testify/assert"
)

func TestHandleError_NotFound(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	notFoundError := fmt.Errorf("%w: about not found", domainerrors.ErrNotFound)

	errorhandler.HandleError(responseRecorder, notFoundError)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"error":"resource not found: about not found"}`, responseRecorder.Body.String())
}

func TestHandleError_ValidationError(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	validationError := fmt.Errorf("%w: label is required", domainerrors.ErrDomainEntityValidation)

	errorhandler.HandleError(responseRecorder, validationError)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"error":"domain entity validation error: label is required"}`, responseRecorder.Body.String())
}

func TestHandleError_InternalServerError(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	unexpectedError := errors.New("database connection failed")

	errorhandler.HandleError(responseRecorder, unexpectedError)

	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"error":"internal server error"}`, responseRecorder.Body.String())
	assert.NotContains(t, responseRecorder.Body.String(), "database connection failed")
}
