package main

import (
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp4", ":0")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Serving %s at http://%s/. Ctrl+C to exit", cwd, l.Addr())
	log.Fatal(http.Serve(l, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Served %s to %s", r.URL, r.RemoteAddr)
			http.FileServer(http.Dir(cwd)).ServeHTTP(w, r)
		})))
}
