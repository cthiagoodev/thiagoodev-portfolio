package handlers

import (
	"net/http"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/usecases"
	errorhandler "github.com/cthiagoodev/portfolio/services/app-api/internal/presentation/error_handler"
	httpresponse "github.com/cthiagoodev/portfolio/services/app-api/pkg/http_response"
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
		errorhandler.HandleError(responseWriter, err)
		return
	}

	httpresponse.WriteJSONResponse(
		responseWriter,
		http.StatusOK,
		about,
	)
}
