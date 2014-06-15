package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func JailControlStatusHandler(res http.ResponseWriter, req *http.Request) {
	fail2banInput := make([]string, 2)
	fail2banInput[0] = "status"
	fail2banInput[1] = mux.Vars(req)["jail"]

	fail2banOutput, err := fail2banRequest(fail2banInput)
	if err != nil {
	}

	//TODO use reflection to assert data structures and give proper errors
	action := fail2banOutput.([]interface{})[1].([]interface{})[1].([]interface{})[1]
	filter := fail2banOutput.([]interface{})[1].([]interface{})[0].([]interface{})[1]

	output := make(map[string]map[string]interface{})
	output["action"] = make(map[string]interface{})
	output["filter"] = make(map[string]interface{})

	output["filter"]["currentlyFailed"] = filter.([]interface{})[0].([]interface{})[1]
	output["filter"]["totalFailed"] = filter.([]interface{})[1].([]interface{})[1]
	output["filter"]["fileList"] = filter.([]interface{})[2].([]interface{})[1]

	output["action"]["currentlyBanned"] = action.([]interface{})[0].([]interface{})[1]
	output["action"]["totalBanned"] = action.([]interface{})[1].([]interface{})[1]
	output["action"]["ipList"] = action.([]interface{})[2].([]interface{})[1]

	encodedOutput, err := json.Marshal(output)
	if err != nil {
	}

	res.Write(encodedOutput)
}

func JailControlHandler(basicRouter *mux.Router) {
	basicRouter.HandleFunc("/status/{jail}", JailControlStatusHandler).Methods("GET")
}
