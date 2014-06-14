package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	GlobalHandler(r.PathPrefix("/global").Subrouter())
	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
