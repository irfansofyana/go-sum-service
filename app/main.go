package main

import (
	"fmt"
	"log"
	"net/http"

	m "sum/middleware"
	"sum/service"
	"sum/utils"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func getSumHandler(w http.ResponseWriter, r *http.Request) {
	var a, _ = r.URL.Query()["a"]
	var b, _ = r.URL.Query()["b"]

	result, err := service.Sum(a[0], b[0])

	if err != nil {
		message := fmt.Sprintf("Internal server error: %s", err)
		utils.GenerateJSONResponse(w, r, http.StatusInternalServerError, message, nil)
		return
	}

	message := "Success"
	utils.GenerateJSONResponse(w, r, http.StatusOK, message, result)
}

func middlewareConfig() *negroni.Negroni {
	router := mux.NewRouter()
	router.Methods("GET").
		Path("/sum").
		HandlerFunc(getSumHandler)

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

	return n
}

func main() {
	port := "8080"

	fmt.Printf("Sum service started at localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), middlewareConfig()))
}
