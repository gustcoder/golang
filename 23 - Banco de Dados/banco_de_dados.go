package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // import é implícito pois o arquivo nao utiliza o driver, mas o pacote "database/sql" sim
)

func main() {
	mysqlConnection := "golang@localhost:golang@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", mysqlConnection)

	// aqui so testa erros anormais de conexao. se usuario/senha estiverem errados, não vai gerar erro aqui
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	defer db.Close()

	// db.Ping() testa de fato a conexao, usuario/senha etc
	if erro := db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Connection is open...")

	users, errorUsers := db.Query("select * from usuarios")

	if errorUsers != nil {
		log.Fatal(errorUsers)
	}
	defer users.Close() // fechar resultado da query como se fosse conexao
	fmt.Println(users)

}
