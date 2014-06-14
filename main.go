package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Configuration struct {
	Addr string
}

func main() {
	file, fileErr := os.Open("config.json")

	if fileErr != nil {
		fmt.Println("failed to open config:", fileErr)
		os.Exit(1)
	}

	configuration := new(Configuration)
	configErr := json.NewDecoder(file).Decode(configuration)

	if configErr != nil {
		fmt.Println("config error:", configErr)
		os.Exit(1)
	}

	r := mux.NewRouter()
	GlobalHandler(r.PathPrefix("/global").Subrouter())
	http.Handle("/", r)
	http.ListenAndServe(configuration.Addr, nil)
}
