package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/api/presentation"
	"go-rest-api/domain"
	"go-rest-api/service/produto"
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

func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("teste")
}

func AtualizarProduto(w http.ResponseWriter, r *http.Request) {

}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {

}
