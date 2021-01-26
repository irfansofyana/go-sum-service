package middleware

import (
	"fmt"
	"net/http"
)

func GenerateJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf("{message: %q}", message)))
}