package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/usecases"
)

type AboutHandler struct {
	getAboutUseCase usecases.GetAboutUseCase
}

func NewAboutHandler(getAboutUseCase usecases.GetAboutUseCase) *AboutHandler {
	return &AboutHandler{
		getAboutUseCase: getAboutUseCase,
	}
}

func (handler *AboutHandler) Get(responseWriter http.ResponseWriter, request *http.Request) {
	about, err := handler.getAboutUseCase.Execute(request.Context())
	if err != nil {
		if errors.Is(err, domainerrors.ErrNotFound) {
			responseWriter.Header().Set("Content-Type", "application/json")
			responseWriter.WriteHeader(http.StatusNotFound)
			_, _ = responseWriter.Write([]byte(`{"error": "about not found"}`))
			return
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, _ = responseWriter.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(responseWriter).Encode(about)
}
