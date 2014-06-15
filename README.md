# fail2rest

##Overview

fail2rest is a small REST server that aims to replicate the failban-client UI

fail2rest will eventually be used as a backend to a small web app to make fail2ban
administration and reporting easier.


##Contributing
Every PR will be merged! Feel free to open up PRs that aren't fully done, I will do
my best to finish them for you. I will make sure to review everything I can. If
you are interested in working on fail2rest, but don't know where to start here are some ideas.

* Find unimplemented fail2ban-client commands
* Improve data assertions before json.marshall (this is really important!)
* Document current API calls (and examples with cURL), small static website for this info
* Start on fail2web, I would like to write this in angular, angular-ui, browserify and d3.js (Open to suggestions just most comfortable with this)
* Write tests, and implement some post-commit system for running tests
* Expand fail2ban-server so that we can perform more operations via socket. I would like to avoid editing files as long as possible

##License
iThe MIT License (MIT)

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
