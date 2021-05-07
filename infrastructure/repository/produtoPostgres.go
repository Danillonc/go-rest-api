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

func (r *PostgresSQL) CriarProduto(produto *domain.Produto) {

}

func (r *PostgresSQL) BuscarProduto(idProduto string) *domain.Produto {
	return nil
}

func (r *PostgresSQL) AtualizarProduto(produto *domain.Produto) *domain.Produto {
	return nil
}

func (r *PostgresSQL) DeletarProduto(id string) {

}
