package repositorios

import (
	"database/sql"
	"fmt"
	"inventory/src/modelos"
)

// Produtos representa um repositório de produtos
type Produtos struct {
	db *sql.DB
}

// NovoRepositorioDeProdutos cria um repositório de produtos
func NovoRepositorioDeProdutos(db *sql.DB) *Produtos {
	return &Produtos{db}
}

// Criar insere um produto no banco de dados
func (repositorio Produtos) Criar(produto modelos.Produto) error {
	statement, erro := repositorio.db.Prepare(
		"insert into inventory (nome, preco, quantidade) values ($1, $2, $3);",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(produto.Nome, produto.Preco, produto.Quantidade)
	if erro != nil {
		return erro
	}

	return nil
}

// Buscar traz todos os produtos que atendem um filtro de nome
func (repositorio Produtos) Buscar(nome string) ([]modelos.Produto, error) {
	nome = fmt.Sprintf("%%%s%%", nome) // %nome%

	linhas, erro := repositorio.db.Query(
		"select id, nome, preco, quantidade from inventory where nome LIKE $1",
		nome,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var produtos []modelos.Produto

	for linhas.Next() {
		var produto modelos.Produto

		if erro = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Preco,
			&produto.Quantidade,
		); erro != nil {
			return nil, erro
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}

// BuscarPorID traz um produto do banco de dados
func (repositorio Produtos) BuscarPorID(ID uint64) (modelos.Produto, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, preco, quantidade from inventory where id = $1",
		ID,
	)
	if erro != nil {
		return modelos.Produto{}, erro
	}
	defer linhas.Close()

	var produto modelos.Produto

	if linhas.Next() {
		if erro = linhas.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Preco,
			&produto.Quantidade,
		); erro != nil {
			return modelos.Produto{}, erro
		}
	}

	return produto, nil
}

// Atualizar altera as informações de um produto no banco de dados
func (repositorio Produtos) Atualizar(ID uint64, produto modelos.Produto) error {
	// var produto_preco = produto.Preco

	var query = "update inventory set nome = $1, preco = $2, quantidade = $3 WHERE id = $4"
	statement, erro := repositorio.db.Prepare(
		query,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(produto.Nome, produto.Preco, produto.Quantidade, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um produto no banco de dados
func (repositorio Produtos) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from inventory where id = $1")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
