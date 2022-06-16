package main

import (
	"net/http"
	"text/template"

	"database/sql"

	_ "github.com/lib/pq"
)

type Product struct {
	Name,
	Description string
	Price    float64
	Quantity int
}

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDB()

	defer db.Close()

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func connectDB() *sql.DB {
	connection := "user=leonardo-dev dbname=go_store password=postgres host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func index(writer http.ResponseWriter, request *http.Request) {
	products := []Product{
		{"RTX 2060", "Muito potente", 1000, 5},
		{"Ryzen 5 5600X", "Ta maluko", 500, 2},
	}

	htmlTemplates.ExecuteTemplate(writer, "Index", products)
}
