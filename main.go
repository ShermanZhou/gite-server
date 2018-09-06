package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	PORT := flag.Int("port", 8080, "listening port")
	HOST := flag.String("host", "127.0.0.1", "binding host")
	BASEDIR := flag.String("path", ".", "base dir of static files")
	BASEURL := flag.String("baseurl", "/", "base url")
	flag.Parse()

	fs := http.FileServer(http.Dir(*BASEDIR))
	http.Handle(*BASEURL, fs)
	fmt.Printf("\nListening on %s:%d...\n", *HOST, *PORT)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *HOST, *PORT), fs)
}
