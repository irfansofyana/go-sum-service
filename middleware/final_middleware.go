package middleware

import (
	"net/http"
	"sum/utils"
)

// FinalMiddleware is the last middleware that will be executed
func FinalMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	utils.GenerateJSONResponse(w, r, 404, "Resource not found", nil)
}
