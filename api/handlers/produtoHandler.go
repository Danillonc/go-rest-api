package handlers

import (
	"encoding/json"
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

func ListarProdutos(service *produto.ProdutoService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		produtosDomain := service.ListarProdutosService()
		pPresentation := presentation.Produto{}
		produtosPresentation := []presentation.Produto{}

		for i := range produtosDomain {
			pPresentation.Id = produtosDomain[i].Id
			pPresentation.Nome = produtosDomain[i].Nome
			pPresentation.Descricao = produtosDomain[i].Descricao
			pPresentation.Preco = produtosDomain[i].Preco
			pPresentation.Quantidade = produtosDomain[i].Quantidade

			produtosPresentation = append(produtosPresentation, pPresentation)
		}
		json.NewEncoder(w).Encode(produtosPresentation)
	})
}

func AtualizarProduto(w http.ResponseWriter, r *http.Request) {

}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {

}
