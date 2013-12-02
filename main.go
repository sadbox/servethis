package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var port string
	var directory string
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&port, "port", "9001", "Port that you want the files served on")
	flag.StringVar(&directory, "directory", cwd, "Directory that you wish to be served")
	flag.Parse()

	log.Printf("Serving %s at http://localhost:%s/. Ctrl+C to exit", directory, port)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("."))))
}
