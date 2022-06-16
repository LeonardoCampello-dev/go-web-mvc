package main

import (
	"net/http"
	"text/template"

	"database/sql"

	_ "github.com/lib/pq"
)

type Product struct {
	Id int
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

func connectDB() *sql.DB {
	connection := "user=leonardo-dev dbname=go_store password=postgres host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func index(writer http.ResponseWriter, request *http.Request) {
	db := connectDB()

	selectAllProducts, err := db.Query("select * from products")

	if err != nil {
		panic(err.Error())
	}

	currentProduct := Product{}

	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		currentProduct.Name = name
		currentProduct.Description = description
		currentProduct.Price = price
		currentProduct.Quantity = quantity

		products = append(products, currentProduct)
	}

	htmlTemplates.ExecuteTemplate(writer, "Index", products)

	defer db.Close()
}
