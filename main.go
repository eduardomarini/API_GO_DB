package main

import (
	"fmt"
	"log"

	"github.com/eduardomarini/API_GO_DB/db"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	database, err := db.Connect()
	if err != nil {
		// Loga o erro e encerra o programa se a conexão falhar
		log.Fatal(err)
	}

	defer func() {
		if err := db.Close(database); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Conexão com o banco de dados bem-sucedida!")
}
