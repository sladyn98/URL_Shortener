package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"


	urlshort "github.com/sladyn98/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Executing the yaml handler")
	dir, err := os.Getwd()
	fmt.Println("Reading the file",dir+"/urls.yml")
	data, err := ioutil.ReadFile(dir+"/urls.yml")
		if err != nil {
			fmt.Println(err)
		}
	yamlHandler, err := urlshort.YAMLHandler(data, mapHandler)
	if err != nil {
		panic(err)
	}
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
