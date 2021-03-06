package middleware

import (
	"net/http"
	"sum/utils"
)

// SumMiddleware is a middleware for /sum endpoint
func SumMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var a, isExistA = r.URL.Query()["a"]
	var b, isExistB = r.URL.Query()["b"]

	if !isExistA || !isExistB {
		message := "Variable a and/or b is not available on query parameter"
		utils.GenerateJSONResponse(w, r, http.StatusBadRequest, message, nil)
		return
	}

	if len(a) != 1 || len(b) != 1 {
		message := "Ambigous value on variable a and/or b"
		utils.GenerateJSONResponse(w, r, http.StatusBadRequest, message, nil)
		return
	}

	if !utils.IsInt(a[0]) || !utils.IsInt(b[0]) {
		message := "Variable a and/or is not a valid integer"
		utils.GenerateJSONResponse(w, r, http.StatusBadRequest, message, nil)
		return
	}

	next(w, r)
}
