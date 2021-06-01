package produto

import "go-rest-api/domain"

type ProdutoRepository interface {
	CriarProduto(produto domain.Produto)
	ListarProdutos() []domain.Produto
	BuscarProduto(id string) *domain.Produto
	AtualizarProduto(produto domain.Produto) *domain.Produto
	DeletarProduto(id string)
}
