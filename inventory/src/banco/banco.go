package banco

import (
	"database/sql"
	"inventory/src/config"

	_ "gorm.io/driver/postgres" //Driver
)

func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("postgres", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
