package common

import (
	"encoding/json"
	"net/http"
)

func HandleError(writer http.ResponseWriter, err error, statusCode *int) {
	encoder := json.NewEncoder(writer)

	if statusCode != nil {
		writer.WriteHeader(*statusCode)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	encoder.Encode(map[string]any{
		"status": statusCode,
		"error":  err.Error(),
	})
}
