package httpresponse_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	httpresponse "github.com/cthiagoodev/portfolio/services/app-api/pkg/http_response"
	"github.com/stretchr/testify/assert"
)

type samplePayload struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func TestWriteJSONResponse_Success_StructPayload(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	payload := samplePayload{
		Name: "Thiago Sousa",
		Role: "Software Engineer",
	}

	httpresponse.WriteJSONResponse(responseRecorder, http.StatusOK, payload)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"name":"Thiago Sousa","role":"Software Engineer"}`, responseRecorder.Body.String())
}

func TestWriteJSONResponse_Success_EmptySlicePayload(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	payload := []samplePayload{}

	httpresponse.WriteJSONResponse(responseRecorder, http.StatusOK, payload)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `[]`, responseRecorder.Body.String())
}

func TestWriteJSONResponse_Success_CustomStatusCode(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	payload := map[string]string{
		"status": "created",
	}

	httpresponse.WriteJSONResponse(responseRecorder, http.StatusCreated, payload)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
	assert.Equal(t, "application/json", responseRecorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"status":"created"}`, responseRecorder.Body.String())
}
