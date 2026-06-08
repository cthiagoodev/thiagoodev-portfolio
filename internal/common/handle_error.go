package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

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

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"status": statusCode,
		"error":  err.Error(),
	})
}
