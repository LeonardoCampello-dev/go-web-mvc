package routes

import (
	"net/http"

	"github.com/LeonardoCampello-dev/go-web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create-new-product", controllers.CreateNewProduct)
}
