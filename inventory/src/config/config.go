package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o Postgres
	StringConexaoBanco = ""

	// Porta onde a API vai estar rodando
	Porta = 0
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() string {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("Port"))
	if erro != nil {
		Porta = 8000
	}

	StringConexaoBanco = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("Host"),
		os.Getenv("DbPort"),
		os.Getenv("User"),
		os.Getenv("Password"),
		os.Getenv("DbName"),
	)

	return StringConexaoBanco
}
