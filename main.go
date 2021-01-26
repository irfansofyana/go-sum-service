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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Server Error: Error when using sum function"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{sum: %s}", result)))
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