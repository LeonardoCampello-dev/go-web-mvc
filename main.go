package main

import (
	"net/http"

	"github.com/LeonardoCampello-dev/go-web/routes"
)

func main() {
	routes.LoadRoutes()

	http.ListenAndServe(":8000", nil)
}
