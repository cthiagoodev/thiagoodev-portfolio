package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/usecases"
)

type SocialContactHandler struct {
	getAllSocialContactsUseCase usecases.GetAllSocialContactsUseCase
}

func NewSocialContactHandler(getAllSocialContactsUseCase usecases.GetAllSocialContactsUseCase) *SocialContactHandler {
	return &SocialContactHandler{
		getAllSocialContactsUseCase: getAllSocialContactsUseCase,
	}
}

func (handler *SocialContactHandler) GetAll(responseWriter http.ResponseWriter, request *http.Request) {
	socialContacts, err := handler.getAllSocialContactsUseCase.Execute(request.Context())
	if err != nil {
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, _ = responseWriter.Write([]byte(`{"error": "internal server error"}`))
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(responseWriter).Encode(socialContacts)
}

