package common

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, err string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)

	encoder.Encode(map[string]any{
		"status": statusCode,
		"error":  err,
	})
}
