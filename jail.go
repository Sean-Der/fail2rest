package main

import (
	"encoding/json"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
)

func jailGetHandler(res http.ResponseWriter, req *http.Request) {
	jailStatus, _ := fail2go.JailStatus(mux.Vars(req)["jail"])
	jailFailRegex, _ := fail2go.JailFailRegex(mux.Vars(req)["jail"])

	output := make(map[string]interface{})

	for key, value := range jailStatus {
		output[key] = value
	}
	for key, value := range jailFailRegex {
		output[key] = value
	}

	encodedOutput, err := json.Marshal(output)
	if err != nil {
	}

	res.Write(encodedOutput)
}

type jailBanIPBody struct {
	IP string
}

func jailBanIPHandler(res http.ResponseWriter, req *http.Request) {
	var input jailBanIPBody
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, _ := fail2go.JailBanIP(mux.Vars(req)["jail"], input.IP)

	encodedOutput, err := json.Marshal(output)
	if err != nil {
	}

	res.Write(encodedOutput)
}

func jailUnbanIPHandler(res http.ResponseWriter, req *http.Request) {
	var input jailBanIPBody
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, _ := fail2go.JailUnbanIP(mux.Vars(req)["jail"], input.IP)

	encodedOutput, err := json.Marshal(output)
	if err != nil {
	}

	res.Write(encodedOutput)

}

func jailHandler(jailRouter *mux.Router) {
	jailRouter.HandleFunc("/{jail}/banip", jailBanIPHandler).Methods("POST")
	jailRouter.HandleFunc("/{jail}/unbanip", jailUnbanIPHandler).Methods("POST")

	jailRouter.HandleFunc("/{jail}", jailGetHandler).Methods("GET")

}
