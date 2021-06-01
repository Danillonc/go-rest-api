package repository

import (
	"database/sql"
	"go-rest-api/domain"
)

type PostgresSQL struct {
	db *sql.DB
}

func NewPostgresSQL(db *sql.DB) *PostgresSQL {
	return &PostgresSQL{
		db: db,
	}
}

func (r *PostgresSQL) CriarProduto(produto domain.Produto) {
	//camada de criação de dados na base mysql.
	inserirProduto, err := r.db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	inserirProduto.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)
	// defer r.db.Close()

}

func (r *PostgresSQL) ListarProdutos() []domain.Produto {
	produtosDb, err := r.db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := domain.Produto{}
	produtos := []domain.Produto{}

	for produtosDb.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := produtosDb.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	return produtos
}

func BuscarProduto(id string) *domain.Produto {

}

func (r *PostgresSQL) AtualizarProduto(produto domain.Produto) *domain.Produto {
	return nil
}

func (r *PostgresSQL) DeletarProduto(id string) {

}
