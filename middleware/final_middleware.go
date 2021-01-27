package middleware

import (
	"net/http"
	"sum/utils"
)

func FinalMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	utils.GenerateJSONResponse(w, r, 404, "Resource not found", nil)
}