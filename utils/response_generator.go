package utils

import (
	"fmt"
	"math/big"
	"net/http"
)

// GenerateJSONResponse is a function to generate JSON response of the API
func GenerateJSONResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string, sum *big.Int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var response string = ""
	if sum == nil {
		response = fmt.Sprintf("{message:%q, sum:null}", message)
	} else {
		response = fmt.Sprintf("{message:%q, sum:%v}", message, sum)
	}
	w.Write([]byte(response))
}
