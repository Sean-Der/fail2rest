default: run

run:
	go run *.go

libs:
	go get github.com/kisielk/og-rek
	go get github.com/gorilla/mux
