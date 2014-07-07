package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorBody struct {
	Error string
}

func writeHTTPError(res http.ResponseWriter, err error) {
	res.WriteHeader(400)
	encodedOutput, err := json.Marshal(ErrorBody{Error: err.Error()})
	if err != nil {
		fmt.Println("Failed to generate HTTP error: " + err.Error())
	} else {
		res.Write(encodedOutput)
	}

}
