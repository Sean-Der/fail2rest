package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sean-der/fail2go"
)

type Configuration struct {
	Addr           string
	Fail2banSocket string
}

var fail2goConn *fail2go.Conn

func main() {
	configPath := flag.String("config", "config.json", "path to config.json")
	flag.Parse()

	file, fileErr := os.Open(*configPath)

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
	r.HandleFunc("/whois/{object}", func(res http.ResponseWriter, req *http.Request) {
		whoisHandler(res, req, fail2goConn)
	}).Methods("GET")

	http.Handle("/", r)
	fmt.Println(http.ListenAndServe(configuration.Addr, nil))
}
