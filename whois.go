package main

import (
	"encoding/json"
	"github.com/Sean-Der/fail2go"
	"github.com/Sean-Der/goWHOIS"
	"github.com/gorilla/mux"
	"net/http"
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
