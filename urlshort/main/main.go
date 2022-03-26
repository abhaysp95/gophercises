package main

import (
	"fmt"
	"gophercises/urlshort"
	"net/http"
	"os"
)

func main() {
	mux := defaultMux()

	/* pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux) */

	yaml := `
- path: /urlshort
  url: https://godoc.org/github.com/gophercises/urlshort
- path: /yaml-godoc
  url: https://godoc.org/gopkg.in/yaml.v2
`

	yamlHandler, err := urlshort.YamlHandler([]byte(yaml), mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Starting server on 8080")
	http.ListenAndServe("localhost:8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello")
}
