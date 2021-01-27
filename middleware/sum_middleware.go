package middleware

import (
	"net/http"
	"sum/utils"
)

// func SumMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
// 		var a, isExistA = r.URL.Query()["a"]
// 		var b, isExistB = r.URL.Query()["b"]

// 		if (!isExistA || !isExistB) {	
// 			message := "Variable a and/or b is not available on query parameter"
// 			utils.GenerateJSONResponse(w, r, 400, message, "")
// 			return
// 		}

// 		if (len(a) != 1 || len(b) != 1) {
// 			message := "Ambigous value on variable a and/or b"
// 			utils.GenerateJSONResponse(w, r, 400, message, "")
// 			return
// 		}

// 		if (!utils.IsInt(a[0]) || !utils.IsInt(b[0])) {
// 			message := "Variable a and/or is not a valid integer"
// 			utils.GenerateJSONResponse(w, r, 400, message, "")
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func GetSumHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var a, isExistA = r.URL.Query()["a"]
	var b, isExistB = r.URL.Query()["b"]

	if (!isExistA || !isExistB) {	
		message := "Variable a and/or b is not available on query parameter"
		utils.GenerateJSONResponse(w, r, 400, message, "")
		return
	}

	if (len(a) != 1 || len(b) != 1) {
		message := "Ambigous value on variable a and/or b"
		utils.GenerateJSONResponse(w, r, 400, message, "")
		return
	}

	if (!utils.IsInt(a[0]) || !utils.IsInt(b[0])) {
		message := "Variable a and/or is not a valid integer"
		utils.GenerateJSONResponse(w, r, 400, message, "")
		return
	}

	next(w, r)
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	next(rw, r)
	// do some stuff after
}