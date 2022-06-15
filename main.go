package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name,
	Description string
	Price    float64
	Quantity int
}

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	products := []Product{
		{"RTX 2060", "Muito potente", 1000, 5},
		{"Ryzen 5 5600X", "Ta maluko", 500, 2},
	}

	htmlTemplates.ExecuteTemplate(writer, "Index", products)
}
