package main

import (
	"go-rest-api/routes"
	"net/http"
)

func main() {
	routes.LoadApiRoutes()
	http.ListenAndServe(":8080", nil)
}
