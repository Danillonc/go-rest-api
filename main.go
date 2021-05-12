package main

import (
	"go-rest-api/config/db"
	"go-rest-api/infrastructure/repository"
	"go-rest-api/routes"
	"go-rest-api/service/produto"
	"log"
	"net/http"
)

func main() {
	db := db.OpenConnectionDb()
	produtoRepo := repository.NewPostgresSQL(db)
	produtoService := produto.NewService(produtoRepo)
	routes := routes.LoadApiRoutes(produtoService)

	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal(err.Error())
	}

}
