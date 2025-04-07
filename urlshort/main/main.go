package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/0xgouda/urlshort"
)

func main() {
	ymlFile := flag.String("Yaml File", "urls.yml", "Yaml File Containing path & url redirection pairs")
	flag.Parse()

	yaml, err := os.ReadFile(*ymlFile)
	urlshort.CheckErr(err)

	mux := defaultMux()
	yamlHandler, err := urlshort.YAMLHandler(yaml, mux)
	urlshort.CheckErr(err)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
