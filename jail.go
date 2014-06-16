package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/Sean-Der/fail2go"
	"net/http"
)

func jailGetHandler(res http.ResponseWriter, req *http.Request) {
	jailStatus, _ := fail2go.JailStatus(mux.Vars(req)["jail"])

	output := make(map[string]interface{})

	for key, value := range jailStatus {
		output[key] = value
	}
	encodedOutput, err := json.Marshal(output)
	if err != nil {
	}

	res.Write(encodedOutput)
}

func jailHandler(jailRouter *mux.Router) {
	jailRouter.HandleFunc("/{jail}", jailGetHandler).Methods("GET")
}
