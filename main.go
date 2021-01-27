package main

import (
	"fmt"
	"log"
	"net/http"
	
	"sum/utils"
	m"sum/middleware"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func GetSumHandler(w http.ResponseWriter, r *http.Request) {
	var a, _ = r.URL.Query()["a"]
	var b, _ = r.URL.Query()["b"]

	result, err := utils.Sum(a[0], b[0])
	
	if (err != nil) {
		message := "Internal server error"
		utils.GenerateJSONResponse(w, r, 500, message, "")
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

	n := negroni.New()
	n.Use(negroni.HandlerFunc(m.GetSumHandler))
	n.UseHandler(router)
	
	log.Fatal(http.ListenAndServe(":8080", n))
}