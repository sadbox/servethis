servethis
=========

Simple golang-powered file server

usage
=====
Run ```servethis``` by itself to have it serve the local directory on http://localhost:9001/

You can also specify the port and directory to be served by using the -p and -d flags.

example
=======
```
jmcguire@farnsworth:~/go/src/github.com/sadbox/servethis$ ./servethis 
2013/12/01 22:37:10 Serving /home/jmcguire/go/src/github.com/sadbox/servethis at http://localhost:9001/. Ctrl+C to exit
2013/12/01 22:37:12 Served /.gitignore to [::1]:34231
2013/12/01 22:37:27 Served / to [::1]:34231
```

install
=======
```go get github.com/sadbox/servethis```
