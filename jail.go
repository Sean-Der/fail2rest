package main

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sean-der/fail2go"
)

func jailGetHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	currentlyFailed, totalFailed, fileList, currentlyBanned, totalBanned, IPList, err := fail2goConn.JailStatus(mux.Vars(req)["jail"])
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	failRegexes, _ := fail2goConn.JailFailRegex(mux.Vars(req)["jail"])
	findTime, _ := fail2goConn.JailFindTime(mux.Vars(req)["jail"])
	useDNS, _ := fail2goConn.JailUseDNS(mux.Vars(req)["jail"])
	maxRetry, _ := fail2goConn.JailMaxRetry(mux.Vars(req)["jail"])
	actions, _ := fail2goConn.JailActions(mux.Vars(req)["jail"])

	if IPList == nil {
		IPList = []string{}
	}
	if failRegexes == nil {
		failRegexes = []string{}
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{
		"currentlyFailed": currentlyFailed,
		"totalFailed":     totalFailed,
		"fileList":        fileList,
		"currentlyBanned": currentlyBanned,
		"totalBanned":     totalBanned,
		"IPList":          IPList,
		"failRegexes":     failRegexes,
		"findTime":        findTime,
		"useDNS":          useDNS,
		"maxRetry":        maxRetry,
		"actions":         actions})
	res.Write(encodedOutput)
}

type jailBanIPBody struct {
	IP string
}

func jailBanIPHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailBanIPBody
	json.NewDecoder(req.Body).Decode(&input)

	output, err := fail2goConn.JailBanIP(mux.Vars(req)["jail"], input.IP)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"bannedIP": output})
	res.Write(encodedOutput)
}

func jailUnbanIPHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailBanIPBody
	json.NewDecoder(req.Body).Decode(&input)
	output, err := fail2goConn.JailUnbanIP(mux.Vars(req)["jail"], input.IP)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"unBannedIP": output})
	res.Write(encodedOutput)
}

type jailFailRegexBody struct {
	FailRegex string
}

func jailAddFailRegexHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFailRegexBody
	json.NewDecoder(req.Body).Decode(&input)

	output, err := fail2goConn.JailAddFailRegex(mux.Vars(req)["jail"], input.FailRegex)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"FailRegex": output})
	res.Write(encodedOutput)
}

func jailDeleteFailRegexHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFailRegexBody
	json.NewDecoder(req.Body).Decode(&input)

	output, err := fail2goConn.JailDeleteFailRegex(mux.Vars(req)["jail"], input.FailRegex)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"FailRegex": output})
	res.Write(encodedOutput)
}

type RegexResult struct {
	Line  string
	Match bool
}

func jailTestFailRegexHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFailRegexBody
	json.NewDecoder(req.Body).Decode(&input)

	regexp, err := regexp.Compile(strings.Replace(input.FailRegex, "<HOST>", "(?:::f{4,6}:)?(?P<host>\\S+)", -1))

	if err != nil {
		writeHTTPError(res, err)
		return
	}

	_, _, fileList, _, _, _, err := fail2goConn.JailStatus(mux.Vars(req)["jail"])
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	output := make(map[string][]RegexResult)
	for _, fileName := range fileList {
		file, err := os.Open(fileName)
		if err != nil {
			writeHTTPError(res, err)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			output[fileName] = append(output[fileName], RegexResult{Match: regexp.MatchString(scanner.Text()), Line: scanner.Text()})
		}
	}

	encodedOutput, _ := json.Marshal(output)
	res.Write(encodedOutput)
}

type jailFindTimeBody struct {
	FindTime int
}

func jailSetFindTimeHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailFindTimeBody
	json.NewDecoder(req.Body).Decode(&input)

	output, err := fail2goConn.JailSetFindTime(mux.Vars(req)["jail"], input.FindTime)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"FindTime": output})
	res.Write(encodedOutput)
}

type jailUseDNSBody struct {
	UseDNS string
}

func jailSetUseDNSHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailUseDNSBody
	json.NewDecoder(req.Body).Decode(&input)

	output, err := fail2goConn.JailSetUseDNS(mux.Vars(req)["jail"], input.UseDNS)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"useDNS": output})
	res.Write(encodedOutput)
}

type jailMaxRetryBody struct {
	MaxRetry int
}

func jailSetMaxRetryHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	var input jailMaxRetryBody
	json.NewDecoder(req.Body).Decode(&input)

	output, err := fail2goConn.JailSetMaxRetry(mux.Vars(req)["jail"], input.MaxRetry)
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{"maxRetry": output})
	res.Write(encodedOutput)
}

func jailActionHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	port, err := fail2goConn.JailActionProperty(mux.Vars(req)["jail"], mux.Vars(req)["action"], "port")
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]interface{}{
		"port": port})
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

	jailRouter.HandleFunc("/{jail}/testfailregex", func(res http.ResponseWriter, req *http.Request) {
		jailTestFailRegexHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}/findtime", func(res http.ResponseWriter, req *http.Request) {
		jailSetFindTimeHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}/usedns", func(res http.ResponseWriter, req *http.Request) {
		jailSetUseDNSHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}/maxretry", func(res http.ResponseWriter, req *http.Request) {
		jailSetMaxRetryHandler(res, req, fail2goConn)
	}).Methods("POST")

	jailRouter.HandleFunc("/{jail}/action/{action}", func(res http.ResponseWriter, req *http.Request) {
		jailActionHandler(res, req, fail2goConn)
	}).Methods("GET")

	jailRouter.HandleFunc("/{jail}", func(res http.ResponseWriter, req *http.Request) {
		jailGetHandler(res, req, fail2goConn)
	}).Methods("GET")
}
