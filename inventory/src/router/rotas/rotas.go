package rotas

import (
	"inventory/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasProdutos
	rotas = append(rotas)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
	}

	return r
}
