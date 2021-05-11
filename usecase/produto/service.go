package produto

import (
	"fmt"
	"go-rest-api/domain"
)

type ProdutoService struct {
	repository ProdutoRepository
}

func NewService(repo ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repo,
	}
}

func (s *ProdutoService) CriarProdutoService(produto *domain.Produto) {
	//chamar repository para persistir os dados.
	s.repository.CriarProduto(produto)
	fmt.Println("produto: ", produto)

}

func (s *ProdutoService) BuscarProdutoService(idProduto string) *domain.Produto {
	return nil

}

func (s *ProdutoService) AtualizarProdutoService(produto *domain.Produto) *domain.Produto {
	return nil

}

func (s *ProdutoService) DeletarProdutoService(id string) {

}
