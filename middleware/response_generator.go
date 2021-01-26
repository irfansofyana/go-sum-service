package middleware

import (
	"fmt"
	"net/http"
)

func GenerateJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string, sum string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	var response string = ""
	if (sum == "" || sum == "null") {
		response = fmt.Sprintf("{message:%q, sum:null}", message)
	} else {
		response = fmt.Sprintf("{message:%q, sum:%q}", message, sum)
	}
	w.Write([]byte(response))
}