package main

import (
	"encoding/json"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
)

func jailGetHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Fail2goConn) {
	currentlyFailed, totalFailed, fileList, currentlyBanned, totalBanned, IPList, _ := fail2goConn.JailStatus(mux.Vars(req)["jail"])
	failRegexes, _ := fail2goConn.JailFailRegex(mux.Vars(req)["jail"])

	encodedOutput, err := json.Marshal(map[string]interface{}{
		"currentlyFailed": currentlyFailed,
		"totalFailed":     totalFailed,
		"fileList":        fileList,
		"currentlyBanned": currentlyBanned,
		"totalBanned":     totalBanned,
		"IPList":          IPList,
		"failregex":       failRegexes})

	if err != nil {
	}

	res.Write(encodedOutput)
}

type jailBanIPBody struct {
	IP string
}

func jailBanIPHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Fail2goConn) {
	var input jailBanIPBody
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, _ := fail2goConn.JailBanIP(mux.Vars(req)["jail"], input.IP)

	encodedOutput, err := json.Marshal(map[string]interface{}{"bannedIP": output})
	if err != nil {
	}

	res.Write(encodedOutput)
}

func jailUnbanIPHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Fail2goConn) {
	var input jailBanIPBody
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, _ := fail2goConn.JailUnbanIP(mux.Vars(req)["jail"], input.IP)

	encodedOutput, err := json.Marshal(map[string]interface{}{"unBannedIP": output})
	if err != nil {
	}

	res.Write(encodedOutput)

}

func jailHandler(jailRouter *mux.Router, fail2goConn *fail2go.Fail2goConn) {

	jailRouter.HandleFunc("/{jail}/bannedips", func(res http.ResponseWriter, req *http.Request) {
		jailBanIPHandler(res, req, fail2goConn)
	}).Methods("POST")
	jailRouter.HandleFunc("/{jail}/bannedips", func(res http.ResponseWriter, req *http.Request) {
		jailUnbanIPHandler(res, req, fail2goConn)
	}).Methods("DELETE")

	jailRouter.HandleFunc("/{jail}", func(res http.ResponseWriter, req *http.Request) {
		jailGetHandler(res, req, fail2goConn)
	}).Methods("GET")
}
