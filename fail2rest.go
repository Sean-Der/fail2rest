package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Configuration struct {
	Addr           string
	Fail2banSocket string
}

var fail2goConn *fail2go.Fail2goConn

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

	fail2goConn := fail2go.Newfail2goConn(configuration.Fail2banSocket)
	r := mux.NewRouter()

	globalHandler(r.PathPrefix("/global").Subrouter(), fail2goConn)
	jailHandler(r.PathPrefix("/jail").Subrouter(), fail2goConn)

	http.Handle("/", r)
	http.ListenAndServe(configuration.Addr, nil)
}
