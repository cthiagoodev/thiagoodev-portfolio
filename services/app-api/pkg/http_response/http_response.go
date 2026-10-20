package httpresponse

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse[T any](responseWriter http.ResponseWriter, statusCode int, data T) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	_ = json.NewEncoder(responseWriter).Encode(data)
}
