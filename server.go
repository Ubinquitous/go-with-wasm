package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir = flag.String("dir", ".", "directory on serve")
)

func main() {
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))));
}