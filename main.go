package main

import (
	"net/http"
	"text/template"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	htmlTemplates.ExecuteTemplate(writer, "Index", nil)
}
