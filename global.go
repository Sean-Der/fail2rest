package main

import (
	"encoding/json"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
)

func globalStatusHandler(res http.ResponseWriter, req *http.Request) {
	globalStatus, _ := fail2go.GlobalStatus()

	encodedOutput, err := json.Marshal(globalStatus)
	if err != nil {
	}

	res.Write(encodedOutput)
}

func globalPingHandler(res http.ResponseWriter, req *http.Request) {
	globalPing, _ := fail2go.GlobalPing()

	encodedOutput, err := json.Marshal(globalPing)
	if err != nil {
	}

	res.Write(encodedOutput)

}

func globalHandler(globalRouter *mux.Router) {
	globalRouter.HandleFunc("/status", globalStatusHandler).Methods("GET")
	globalRouter.HandleFunc("/ping", globalPingHandler).Methods("GET")
}
