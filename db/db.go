package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	user     = "postgres"
	password = "marina16"
	dbname   = "postgres"
	host     = "localhost"
	port     = 5432
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	// verifica a conexão ao banco de dados
	db, err := sql.Open("pgx", connStr) // pgx -> drive para o pacote database/sql
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	// verifica se a conexão é bem-sucedida
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // context.Background() -> cria um contexto vazio / context.WithTimeout() -> cria um contexto com um tempo limite
	defer cancel()                                                          // cancela o contexto após a execução da função

	if err := db.PingContext(ctx); err != nil { // db.PingContext() -> verifica se a conexão com o banco de dados está ativa
		return nil, fmt.Errorf("não foi possível alcançar o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados bem-sucedida!")

	return db, nil
}

// Fechar conexão com o banco de dados
func Close(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("erro ao fechar a conexão com o banco de dados: %v", err)
	}
	return nil
}
