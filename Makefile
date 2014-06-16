default: run

run:
	go run *.go

libs:
	go get github.com/gorilla/mux
	go get github.com/Sean-Der/fail2go
