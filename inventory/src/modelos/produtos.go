package modelos

import (
	"errors"
	"strings"
)

// Produto representa um produto cadastrado, ou não, no estoque
type Produto struct {
	ID         uint64  `json:"id,omitempty"`
	Nome       string  `json:"nome,omitempty"`
	Preco      float64 `json:"preco,omitempty"`
	Quantidade uint64  `json:"quantidade,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o produto recebido
func (produto *Produto) Preparar(etapa string) error {
	if erro := produto.validar(etapa); erro != nil {
		return erro
	}

	if erro := produto.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

// Etapas de erro de preenchimento dos dados
func (produto *Produto) validar(etapa string) error {
	if produto.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if produto.Preco == 0 {
		return errors.New("o preço é obrigatório e não pode estar em branco")
	}

	if produto.Quantidade == 0 {
		return errors.New("a quantidade é obrigatória e não pode estar em branco")
	}

	return nil
}

// Para tirar espaços em branco
func (produto *Produto) formatar(etapa string) error {
	produto.Nome = strings.TrimSpace(produto.Nome)
	return nil
}
