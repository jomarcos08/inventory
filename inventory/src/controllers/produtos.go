package controllers

import (
	"encoding/json"
	"inventory/src/banco"
	"inventory/src/modelos"
	"inventory/src/repositorios"
	"inventory/src/respostas"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarProduto insere um produto no banco de dados
func CriarProduto(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var produto modelos.Produto
	if erro = json.Unmarshal(corpoRequest, &produto); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = produto.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	erro = repositorio.Criar(produto)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, produto)
}

// BuscarProdutos busca todos os produtos salvos no banco
func BuscarProdutos(w http.ResponseWriter, r *http.Request) {

	nome := strings.ToLower(r.URL.Query().Get("produto"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produto, erro := repositorio.Buscar(nome)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, produto)
}

// BuscarProduto busca um produto salvo no banco
func BuscarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	produtoID, erro := strconv.ParseUint(parametros["produtoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produto, erro := repositorio.BuscarPorID(produtoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, produto)
}

// AtualizarProduto altera as informações de um produto no banco
func AtualizarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	produtoID, erro := strconv.ParseUint(parametros["produtoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var produto modelos.Produto
	if erro = json.Unmarshal(corpoRequisicao, &produto); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = produto.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	if erro = repositorio.Atualizar(produtoID, produto); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarProduto exclui as informações de um produto no banco
func DeletarProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	produtoID, erro := strconv.ParseUint(parametros["produtoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	if erro = repositorio.Deletar(produtoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
