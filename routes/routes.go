package routes

import (
	"net/http"

	"github.com/LeonardoCampello-dev/go-web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/insert-new-product-form", controllers.InsertNewProductForm)
	http.HandleFunc("/insert", controllers.InsertNewProduct)
	http.HandleFunc("/delete", controllers.DeleteProduct)
	http.HandleFunc("/edit", controllers.EditProductForm)
	http.HandleFunc("/update", controllers.EditProduct)
}
