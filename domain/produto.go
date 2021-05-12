package domain

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Quantidade int
	Preco      float64
}

//Create a simple product object.
func CriarProduto(nome, descricao string, quantidade int, preco float64) Produto {
	produto := Produto{
		Nome:       nome,
		Descricao:  descricao,
		Quantidade: quantidade,
		Preco:      preco,
	}
	return produto
}
