package main

import (
	"fmt"
	"inventory/src/config"
	"inventory/src/router"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d\n", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
