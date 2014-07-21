default: run

run:
	go run *.go

build:
	go build *.go

libs:
	go get github.com/gorilla/mux
	go get github.com/Sean-Der/fail2go
	go get github.com/Sean-Der/goWHOIS
