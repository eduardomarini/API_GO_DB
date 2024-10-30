package main

import (
	"fmt"
	"log"

	"github.com/eduardomarini/API_GO_DB/db"
	"github.com/eduardomarini/API_GO_DB/migrations"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// Cria uma instância do banco de dados
	database, err := db.Connect()
	if err != nil {
		// Loga o erro e encerra o programa se a conexão falhar
		log.Fatal(err)
	}

	// Faz todas as migrações
	err = migrations.RunMigrations(database)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabela Autores migrada com sucesso!")
}
