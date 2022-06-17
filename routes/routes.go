package routes

import (
	"net/http"

	"github.com/LeonardoCampello-dev/go-web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
