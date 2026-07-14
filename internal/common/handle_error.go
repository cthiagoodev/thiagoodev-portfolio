package common

import (
	"errors"
	"net/http"

	"github.com/cthiagoodev/thiagoodev-portfolio/internal/common/templates"
)

func HandleError(w http.ResponseWriter, tm *templates.TemplateManager, err error) {
	statusCode := http.StatusInternalServerError

	switch {
	case errors.Is(err, ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, ErrInvalidData):
		statusCode = http.StatusBadRequest
	case errors.Is(err, ErrConnection):
		statusCode = http.StatusServiceUnavailable
	default:
		statusCode = http.StatusInternalServerError
	}

	tm.Render(w, "error", statusCode)
}
