package routes

import (
	"go-rest-api/api/handlers"
	"go-rest-api/service/produto"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//Function responsible to configure routes to API REST
func LoadApiRoutes(service *produto.ProdutoService) *mux.Router {
	n := negroni.New()
	route := mux.NewRouter()
	route.Handle("/v1/produto/criar", n.With(negroni.Wrap(handlers.CriarProduto(service)))).Methods("POST")
	// route.HandleFunc("v1/produto/criar", n.With(negroni.Wrap(handlers.CriarProduto(service)))).Methods("POST")
	// route.HandleFunc("v1/produto/buscar", handlers.BuscarProduto).Methods("GET")
	// route.HandleFunc("v1/produto/atualizar", handlers.AtualizarProduto).Methods("PUT")
	// route.HandleFunc("v1/produto/deletar", handlers.DeletarProduto).Methods("DELETE")
	return route

}
