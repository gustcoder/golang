package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // import é implícito pois o arquivo nao utiliza o driver, mas o pacote "database/sql" sim
)

func Connect() (*sql.DB, error) {
	connection := "golang@localhost:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", connection)

	if err != nil {
		return nil, err
	}

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
