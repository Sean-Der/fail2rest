package main

import (
	"encoding/json"
	"github.com/Sean-Der/fail2go"
	"github.com/gorilla/mux"
	"net/http"
)

func jailGetHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	currentlyFailed, totalFailed, fileList, currentlyBanned, totalBanned, IPList, _ := fail2goConn.JailStatus(mux.Vars(req)["jail"])
	failRegexes, _ := fail2goConn.JailFailRegex(mux.Vars(req)["jail"])
	findTime, _ := fail2goConn.JailFindTime(mux.Vars(req)["jail"])
	useDNS, _ := fail2goConn.JailUseDNS(mux.Vars(req)["jail"])
	maxRetry, _ := fail2goConn.JailMaxRetry(mux.Vars(req)["jail"])

	//If IPList is nil/null/doesn't exist, then initialize it to an empty string array.  This resolves the front end issue where a null value is trying to be parsed for the ips.
	if IPList == nil {
		IPList = []string{}
	}

	encodedOutput, err := json.Marshal(map[string]interface{}{
		"currentlyFailed": currentlyFailed,
		"totalFailed":     totalFailed,
		"fileList":        fileList,
		"currentlyBanned": currentlyBanned,
		"totalBanned":     totalBanned,
		"IPList":          IPList,
		"failRegexes":     failRegexes,
		"findTime":        findTime,
		"useDNS":          useDNS,
		"maxRetry":        maxRetry})

	if err != nil {
	}

	res.Write(encodedOutput)
}

type jailBanIPBody struct {
	IP string
}

func jailBanIPHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
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

func jailUnbanIPHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
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

type jailFailRegexBody struct {
	FailRegex string
}

func jailAddFailRegexHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFailRegexBody
	var encodedOutput []byte

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {

	}

	output, err := fail2goConn.JailAddFailRegex(mux.Vars(req)["jail"], input.FailRegex)
	if err != nil {
		res.WriteHeader(400)
		encodedOutput, err = json.Marshal(ErrorBody{Error: err.Error()})
	} else {
		encodedOutput, err = json.Marshal(map[string]interface{}{"FailRegex": output})
	}

	if err != nil {
	}

	res.Write(encodedOutput)
}

func jailDeleteFailRegexHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFailRegexBody
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, _ := fail2goConn.JailDeleteFailRegex(mux.Vars(req)["jail"], input.FailRegex)

	encodedOutput, err := json.Marshal(map[string]interface{}{"FailRegex": output})
	if err != nil {
	}

	res.Write(encodedOutput)
}

type jailFindTimeBody struct {
	FindTime int
}

func jailSetFindTimeHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFindTimeBody

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, err := fail2goConn.JailSetFindTime(mux.Vars(req)["jail"], input.FindTime)
	if err != nil {
	}

	encodedOutput, err := json.Marshal(map[string]interface{}{"FindTime": output})

	if err != nil {
	}

	res.Write(encodedOutput)
}

type jailUseDNSBody struct {
	UseDNS string
}

func jailSetUseDNSHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailUseDNSBody

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, err := fail2goConn.JailSetUseDNS(mux.Vars(req)["jail"], input.UseDNS)
	if err != nil {
	}

	encodedOutput, err := json.Marshal(map[string]interface{}{"useDNS": output})

	if err != nil {
	}

	res.Write(encodedOutput)
}

type jailMaxRetryBody struct {
	MaxRetry int
}

func jailSetMaxRetryHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailMaxRetryBody

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
	}

	output, err := fail2goConn.JailSetMaxRetry(mux.Vars(req)["jail"], input.MaxRetry)
	if err != nil {
	}

	encodedOutput, err := json.Marshal(map[string]interface{}{"maxRetry": output})

	if err != nil {
	}

	res.Write(encodedOutput)
}

func jailHandler(jailRouter *mux.Router, fail2goConn *fail2go.Conn) {

	jailRouter.HandleFunc("/{jail}/bannedip", func(res http.ResponseWriter, req *http.Request) {
		jailBanIPHandler(res, req, fail2goConn)
	}).Methods("POST")
	jailRouter.HandleFunc("/{jail}/bannedip", func(res http.ResponseWriter, req *http.Request) {
		jailUnbanIPHandler(res, req, fail2goConn)
	}).Methods("DELETE")

	jailRouter.HandleFunc("/{jail}/failregex", func(res http.ResponseWriter, req *http.Request) {
		jailAddFailRegexHandler(res, req, fail2goConn)
	}).Methods("POST")
	jailRouter.HandleFunc("/{jail}/failregex", func(res http.ResponseWriter, req *http.Request) {
		jailDeleteFailRegexHandler(res, req, fail2goConn)
	}).Methods("DELETE")

	jailRouter.HandleFunc("/{jail}/findtime", func(res http.ResponseWriter, req *http.Request) {
		jailSetFindTimeHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}/usedns", func(res http.ResponseWriter, req *http.Request) {
		jailSetUseDNSHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}/maxretry", func(res http.ResponseWriter, req *http.Request) {
		jailSetMaxRetryHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}", func(res http.ResponseWriter, req *http.Request) {
		jailGetHandler(res, req, fail2goConn)
	}).Methods("GET")
}
