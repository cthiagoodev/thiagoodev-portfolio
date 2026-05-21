package common

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error, statusCode *int) {
	encoder := json.NewEncoder(w)

	if statusCode != nil {
		w.WriteHeader(*statusCode)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	encoder.Encode(map[string]any{
		"status": statusCode,
		"error":  err.Error(),
	})
}
