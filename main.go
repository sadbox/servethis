package main

import (
	"log"
    "net"
    "strings"
	"net/http"
	"os"
)

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Served %s to %s", r.URL, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

    l, err := net.Listen("tcp4", ":0")
    if err != nil {
        panic(err)
    }
    addr := strings.Split(l.Addr().String(), ":")

	log.Printf("Serving %s at http://localhost:%s/. Ctrl+C to exit", cwd, addr[len(addr)-1])
	log.Fatal(http.Serve(l, Log(http.FileServer(http.Dir(cwd)))))
}
