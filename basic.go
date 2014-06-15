package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func BasicStatusHandler(res http.ResponseWriter, req *http.Request) {
	fail2banInput := make([]string, 1)
	fail2banInput[0] = "status"

	output, err := fail2banRequest(fail2banInput)
	if err != nil {
	}

	//TODO use reflection to assert data structures and give proper errors
	jails := output.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	jails = strings.Split(jails.(string), ",")

	encodedOutput, err := json.Marshal(jails)
	if err != nil {
	}

	res.Write(encodedOutput)
}

func BasicPingHandler(res http.ResponseWriter, req *http.Request) {
	fail2banInput := make([]string, 1)
	fail2banInput[0] = "ping"

	output, err := fail2banRequest(fail2banInput)
	if err != nil {
	}

	//TODO use reflection to assert data structures and give proper errors
	output = output.([]interface{})[1]

	encodedOutput, err := json.Marshal(output)
	if err != nil {
	}

	res.Write(encodedOutput)
}

func BasicHandler(basicRouter *mux.Router) {
	basicRouter.HandleFunc("/status", BasicStatusHandler).Methods("GET")
	basicRouter.HandleFunc("/ping", BasicPingHandler).Methods("GET")
}
