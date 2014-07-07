package main

import (
	"encoding/json"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
)

func globalStatusHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	globalStatus, err := fail2goConn.GlobalStatus()

	if err != nil {
		writeHTTPError(res, err)
		return
	}
	encodedOutput, err := json.Marshal(globalStatus)

	if err != nil {
		writeHTTPError(res, err)
		return
	}

	res.Write(encodedOutput)
}

func globalPingHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	globalPing, _ := fail2goConn.GlobalPing()

	encodedOutput, err := json.Marshal(globalPing)
	if err != nil {
	}

	res.Write(encodedOutput)

}

func globalHandler(globalRouter *mux.Router, fail2goConn *fail2go.Conn) {
	globalRouter.HandleFunc("/status", func(res http.ResponseWriter, req *http.Request) {
		globalStatusHandler(res, req, fail2goConn)
	}).Methods("GET")
	globalRouter.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		globalPingHandler(res, req, fail2goConn)
	}).Methods("GET")
}
