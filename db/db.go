package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	user     = "postgres"
	password = "marina16"
	dbname   = "postgres"
	host     = "localhost"
	port     = 5432
)

func Connect() (*gorm.DB, error) {

	// obtém o nome do banco de dados a partir da variável de ambiente ou utiliza o padrão "postgres"
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "postgres"
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)

	// verifica a conexão ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("erro ao configurar a conexão: %v", err)
	}

	// Configurações adicionais de conexão
	sqlDB.SetMaxIdleConns(10)           // número máximo de conexões ociosas
	sqlDB.SetMaxOpenConns(100)          // número máximo de conexões abertas
	sqlDB.SetConnMaxLifetime(time.Hour) // tempo máximo de vida de uma conexão

	fmt.Println("Conexão com o banco de dados bem-sucedida!")
	return db, nil
}

// Fechar conexão com o banco de dados
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("erro ao obter a conexão SQL para fechar: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("erro ao fechar a conexão com o banco de dados: %v", err)
	}
	return nil
}
