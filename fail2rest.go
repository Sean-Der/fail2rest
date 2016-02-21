package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Addr           string
	Fail2banSocket string
	ControllerIp   string
}

var fail2goConn *fail2go.Conn

func controllerIpFilterMiddleware(h http.Handler, allowedIpAddress string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestSource := strings.Split(r.RemoteAddr, ":")
		if requestSource[0] != allowedIpAddress {
			http.Error(w, "Not authorized", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

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

	http.Handle("/", controllerIpFilterMiddleware(r, configuration.ControllerIp))
	fmt.Println(http.ListenAndServe(configuration.Addr, nil))
}
