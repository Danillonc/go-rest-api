package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/api/presentation"
	"net/http"
)

func CriarProduto(w http.ResponseWriter, r *http.Request) {
	//criar uma camada de service(interface) que irá receber um repository, que por sua vez receberá um *DbSQL para persistir os dados.
	var produto presentation.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(produto)
	// fmt.Fprintf(w, "%+v", produto)

}

func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("teste")
}

func AtualizarProduto(w http.ResponseWriter, r *http.Request) {

}

func DeletarProduto(w http.ResponseWriter, r *http.Request) {

}
