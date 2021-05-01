package routes

import (
	"go-rest-api/api/handlers"

	"github.com/gorilla/mux"
)

func LoadApiRoutes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("v1/produto/criar", handlers.CriarProduto).Methods("POST")
	route.HandleFunc("v1/produto/buscar", handlers.BuscarProduto).Methods("GET")
	route.HandleFunc("v1/produto/atualizar", handlers.AtualizarProduto).Methods("PUT")
	route.HandleFunc("v1/produto/deletar", handlers.DeletarProduto).Methods("DELETE")
	return route

}
