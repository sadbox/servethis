package main

import (
	"flag"
	"log"
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
	var port string
	var directory string

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&port, "p", "9001", "Port that you want the files served on")
	flag.StringVar(&directory, "d", cwd, "Directory that you wish to be served")
	flag.Parse()

	log.Printf("Serving %s at http://localhost:%s/. Ctrl+C to exit", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, Log(http.FileServer(http.Dir(directory)))))
}
