package controllers

import (
	"html/template"
	"net/http"

	"github.com/LeonardoCampello-dev/go-web/models"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	products := models.SelectAllProducts()

	htmlTemplates.ExecuteTemplate(writer, "Index", products)
}

func CreateNewProduct(writer http.ResponseWriter, request *http.Request) {
	htmlTemplates.ExecuteTemplate(writer, "CreateNewProduct", nil)
}
