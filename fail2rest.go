package main

import (
	"encoding/json"
	"fmt"
	"path"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Configuration struct {
	Addr           string
	Fail2banSocket string
}

type ErrorBody struct {
	Error string
}

var fail2goConn *fail2go.Conn

func main() {

	//Changing path of current running directory to path of running executable for finding the config.json file.
	//This will allow the creation of a linux service to run fail2rest (ie: service fail2rest start/stop/status)
	var filePath string = path.Dir(os.Args[0])
	file, fileErr := os.Open(filePath + "/config.json")

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
	fmt.Println(http.ListenAndServe(configuration.Addr, nil))
}
