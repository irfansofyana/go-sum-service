package main

import (
	"fmt"
	"log"
	"net/http"

	m "sum/middleware"
	"sum/service"
	"sum/utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

// GetSumHandler is handler function for /sum endpoint
func GetSumHandler(w http.ResponseWriter, r *http.Request) {
	var a, _ = r.URL.Query()["a"]
	var b, _ = r.URL.Query()["b"]

	result, err := service.Sum(a[0], b[0])

	if err != nil {
		message := "Internal server error"
		utils.GenerateJSONResponse(w, r, http.StatusInternalServerError, message, nil)
		return
	}

	message := "Success"
	utils.GenerateJSONResponse(w, r, http.StatusOK, message, result)
}

func readEnvVariables() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func middlewareConfig(router *mux.Router) *negroni.Negroni {
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
	readEnvVariables()
	port := viper.GetString("port")

	fmt.Printf("Sum service started at localhost:%s\n", port)

	router := mux.NewRouter()
	router.Methods("GET").
		Path("/sum").
		HandlerFunc(GetSumHandler)

	log.Fatal(http.ListenAndServe(port, middlewareConfig(router)))
}
