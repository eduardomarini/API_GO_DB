package db

import (
	"database/sql"
	"fmt"
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
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

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
