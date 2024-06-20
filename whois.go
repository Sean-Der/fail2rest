package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sean-der/fail2go"
	"github.com/sean-der/goWHOIS"
)

func whoisHandler(res http.ResponseWriter, req *http.Request, fail2goConn *fail2go.Conn) {
	goWHOISReq := goWHOIS.NewReq(mux.Vars(req)["object"])
	WHOIS, err := goWHOISReq.Raw()
	if err != nil {
		writeHTTPError(res, err)
		return
	}

	encodedOutput, _ := json.Marshal(map[string]string{"WHOIS": WHOIS})
	res.Write(encodedOutput)
}
