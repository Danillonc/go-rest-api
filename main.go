package main

import (
	"go-rest-api/routes"
	"net/http"
)

func main() {
	routes := routes.LoadApiRoutes()
	http.ListenAndServe(":8080", routes)
}
