package produto

import (
	"go-rest-api/domain"
)

type ProdutoService struct {
	repository ProdutoRepository
}

//Function responsible to create a instance from ProdutoService injecting a repository object.
func NewService(repo ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repo,
	}
}

//Function responsible to execute a business logic to create a simple product.
func (s *ProdutoService) CriarProdutoService(produto domain.Produto) {
	s.repository.CriarProduto(produto)
}

func (s *ProdutoService) BuscarProdutoService(idProduto string) *domain.Produto {
	return nil

}

func (s *ProdutoService) AtualizarProdutoService(produto domain.Produto) *domain.Produto {
	return nil

}

func (s *ProdutoService) DeletarProdutoService(id string) {

}
