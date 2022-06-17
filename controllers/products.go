package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/LeonardoCampello-dev/go-web/models"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	products := models.SelectAllProducts()

	htmlTemplates.ExecuteTemplate(writer, "Index", products)
}

func InsertNewProductForm(writer http.ResponseWriter, request *http.Request) {
	htmlTemplates.ExecuteTemplate(writer, "InsertNewProductForm", nil)
}

func InsertNewProduct(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		name := request.FormValue("name")
		description := request.FormValue("description")
		price := request.FormValue("price")
		quantity := request.FormValue("quantity")

		formattedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("price conversion error", err)
		}

		formattedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("quantity conversion error", err)
		}

		models.InsertNewProduct(name, description, formattedPrice, formattedQuantity)
	}

	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}
