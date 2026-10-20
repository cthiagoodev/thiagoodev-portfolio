package error

import (
	"net/http"

	domainerrors "github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
	httpresponse "github.com/cthiagoodev/portfolio/services/app-api/pkg/http_response"
)

func HandleError(responseWriter http.ResponseWriter, err error) {
	body := make(map[string]string)
	var code int

	switch err {
	case domainerrors.ErrNotFound:
		body["error"] = "resource not found %w"
		code = http.StatusNotFound
	default:
		body["error"] = "internal server error %w"
		code = http.StatusInternalServerError
	}

	httpresponse.WriteJSONResponse(
		responseWriter,
		code,
		body,
	)
}
