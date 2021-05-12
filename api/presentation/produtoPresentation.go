package presentation

type Produto struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}
