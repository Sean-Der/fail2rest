#fail2rest

##Overview
fail2rest is a small REST server that aims to allow full administration of a fail2ban server via HTTP

fail2rest will eventually be used as a backend to a small web app to make fail2ban
administration and reporting easier.

##Requirements
fail2rest is written in Golang, so it requires a working Golang installation. If you have never used Golang before install it from your
package manager, and then follow [these](http://golang.org/doc/code.html) instructions to setup your enviroment.

Some libraries currently are not building on older versions of Go, I recommend running the newest stable of Golang.

##Installing
Once you have a working Golang installation all you need to do is run

    go get -v github.com/Sean-Der/fail2rest
    go install -v github.com/Sean-Der/fail2rest

fail2rest will now be available as a binary in your GOPATH, you can run 'fail2rest' from the command line, or copy it
somewhere else to make it available to all users.

Use `whereis fail2rest` to locate your fail2rest binary.
Next you need to configure fail2rest, and then finally make it a system service.

##Configuration
fail2rest is configured via config.json, the default is located [here](https://raw.githubusercontent.com/Sean-Der/fail2rest/master/config.json).
To load your config.json you use the --config flag `fail2rest --config=my_config.json`

fail2rest has two options that be configured via config.json
  * **Fail2banSocket** - The path to the fail2ban socket, can usually be found via `grep socket /etc/fail2ban/fail2ban.conf` you also have to run fail2rest as a user who has permissions to use this socket
  * **Addr** - The address that fail2rest is served upon, it is usually best so serve to the loopback, and then allow access via nginx see an example config in the [fail2web](https://github.com/Sean-Der/fail2web) repository

The default configuration file used by init-scripts is `/etc/fail2rest.json`. You should download the config to your /tmp/ dir and modify it to your needs. 

    cd /tmp/
    wget https://raw.githubusercontent.com/Sean-Der/fail2rest/master/config.json
    
Once you finished editing the configuration file, you should move it from /tmp/config.json to /etc/fail2rest.json executing `mv /tmp/config.json /etc/fail2rest.json`

##Running
Once you have a config.json all you need to do is run `fail2rest --config config.json`

However, fail2rest is designed to run as a service, so init scripts are provided that allow easy management of fail2rest. They can be found [here](https://github.com/Sean-Der/fail2rest/tree/master/init-scripts)
Download the appropriate init file your Distribution. You may need to customize your init script to load your config.json, but most scripts default to /etc/fail2rest.json

##License
The MIT License (MIT)

Copyright (c) 2014 Sean DuBois

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
