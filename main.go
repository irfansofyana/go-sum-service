package main

import (
	"fmt"
	"sum/utils"
	"log"
	"net/http"
	m"sum/middleware"
)

func GetSum(w http.ResponseWriter, r *http.Request) {
	var a, _ = r.URL.Query()["a"]
	var b, _ = r.URL.Query()["b"]

	result, err := utils.Sum(a[0], b[0])
	
	w.Header().Set("Content-Type", "application/json")

	if (err != nil) {
		message := "Internal server error"
		m.GenerateJSONResponse(w, r, 500, message)
		return
	}

	message := fmt.Sprintf("{sum: %s}", result
	m.GenerateJSONResponse(w, r, 200, message)
}

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/sum", GetSum)
	
	var handler http.Handler = mux
	handler = m.SumMiddleware(handler)
	
	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = handler

	fmt.Println("Sum service started at localhost:8080")
	log.Fatal(server.ListenAndServe())
}