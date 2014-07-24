#Running
To quickly run fail2rest just execute `go run *.go`

#Guidelines

* Put as much logic as possible into [fail2go](https://github.com/Sean-Der/fail2go)
* Make sure code is properly formated [gofmt](http://blog.golang.org/go-fmt-your-code)
* Make sure you code compiles
* If adding new REST endpoints try to follow the current style

#REST Style
Currently we have three top level endpoints
* /global (Get/Set information relating to fail2ban)
* /jail/{jail} (Get/Set information relating to a single jail)
* /wwhois/{ip} (Run a WHOIS on the given IP)
