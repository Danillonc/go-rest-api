package service

import (
	"go-rest-api/domain"
	"go-rest-api/repository"
)

type ProdutoService struct {
	repository repository.ProdutoRepository
}

func NewService(repo repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repo,
	}
}

func CriarProduto(produto *domain.Produto) {

}

func BuscarProduto(idProduto string) *domain.Produto {
	return nil
}

func AtualizarProduto(produto *domain.Produto) *domain.Produto {
	return nil
}

func DeletarProduto(id string) {

}
