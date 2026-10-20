package handlers

import (
	"net/http"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/usecases"
	errorhandler "github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/error"
	httpresponse "github.com/cthiagoodev/portfolio/services/app-api/pkg/http_response"
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
		errorhandler.HandleError(responseWriter, err)
		return
	}

	httpresponse.WriteJSONResponse(
		responseWriter,
		http.StatusOK,
		socialContacts,
	)
}
