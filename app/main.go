package main

import (
	"fmt"
	"log"
	"net/http"

	m "sum/middleware"
	"sum/utils"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// GetSumHandler is handler function for /sum endpoint
func GetSumHandler(w http.ResponseWriter, r *http.Request) {
	var a, _ = r.URL.Query()["a"]
	var b, _ = r.URL.Query()["b"]

	result, err := utils.Sum(a[0], b[0])

	if err != nil {
		message := "Internal server error"
		utils.GenerateJSONResponse(w, r, 500, message, nil)
		return
	}

	message := "Success"
	utils.GenerateJSONResponse(w, r, 200, message, result)
}

func main() {
	fmt.Println("Sum service started at localhost:8080")

	router := mux.NewRouter()
	router.Methods("GET").
		Path("/sum").
		HandlerFunc(GetSumHandler)

	middlewareHandler := http.NewServeMux()
	middlewareHandler.Handle("/sum", negroni.New(
		negroni.HandlerFunc(m.SumMiddleware),
		negroni.Wrap(router),
	))
	middlewareHandler.Handle("/", negroni.New(
		negroni.HandlerFunc(m.FinalMiddleware),
		negroni.Wrap(router),
	))

	n := negroni.Classic()
	n.UseHandler(middlewareHandler)
	log.Fatal(http.ListenAndServe(":8080", n))
}
