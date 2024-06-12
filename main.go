package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := strings.Replace(r.URL.Path, "/", "", 1)
	if name == "" {
		name = "World"
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("error parsing template: ", err)
	}

	err = t.Execute(w, name)
	if err != nil {
		log.Fatal("error executing template: ", err)
	}
}
