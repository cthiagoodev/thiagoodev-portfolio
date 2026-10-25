package errorhandler

import (
	"errors"
	"net/http"

	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	httpresponse "github.com/cthiagoodev/portfolio/services/app-api/pkg/http_response"
)

func HandleError(responseWriter http.ResponseWriter, err error) {
	body := make(map[string]string)
	var code int

	switch {
	case errors.Is(err, domainerrors.ErrNotFound):
		code = http.StatusNotFound
		body["error"] = err.Error()
	case errors.Is(err, domainerrors.ErrDomainEntityValidation):
		code = http.StatusBadRequest
		body["error"] = err.Error()
	default:
		code = http.StatusInternalServerError
		body["error"] = "internal server error"
	}

	httpresponse.WriteJSONResponse(
		responseWriter,
		code,
		body,
	)
}
