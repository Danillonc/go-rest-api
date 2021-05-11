package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/api/presentation"
	"go-rest-api/domain"
	"go-rest-api/usecase/produto"
	"net/http"
)

func CriarProduto(service *produto.ProdutoService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var produto presentation.Produto
		err := json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		produtoEntity := domain.CriarProduto(produto.Nome, produto.Descricao, produto.Quantidade, produto.Preco)
		service.CriarProdutoService(produtoEntity)
		json.NewEncoder(w).Encode(produto)
	})

}

// func CriarProduto(w http.ResponseWriter, r *http.Request) {
// 	//modificar o parametro de entrada para receber um service e implementar http request/response no corpo de cada handler func.
// 	//criar uma camada de service(interface) que irá receber um repository, que por sua vez receberá um *DbSQL para persistir os dados.
// 	var produto presentation.Produto
// 	err := json.NewDecoder(r.Body).Decode(&produto)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(produto)
// 	// fmt.Fprintf(w, "%+v", produto)

// }

func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("teste")
}

func AtualizarProduto(w http.ResponseWriter, r *http.Request) {

}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {

}
