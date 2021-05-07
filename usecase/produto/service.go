package produto

import "go-rest-api/domain"

type ProdutoService struct {
	repository ProdutoRepository
}

func NewService(repo ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repo,
	}
}

func CriarProdutoService(produto *domain.Produto) {

}

func BuscarProdutoService(idProduto string) *domain.Produto {
	return nil

}

func AtualizarProdutoService(produto *domain.Produto) *domain.Produto {
	return nil

}

func DeletarProdutoService(id string) {

}
