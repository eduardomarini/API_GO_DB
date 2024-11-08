package routes

import (
	"fmt"
	"os"

	"github.com/eduardomarini/API_GO_DB/db"
	"github.com/eduardomarini/API_GO_DB/models"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, error) { // configura uma instância temporária do BD para os teste

	// Define o nome do banco de dados de teste temporariamente
	os.Setenv("DB_NAME", "postgres_test")

	dsn, err := db.Connect()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	if err := dsn.AutoMigrate(&models.Usuarios{}); err != nil {
		return nil, fmt.Errorf("erro ao migrar banco de dados de teste: %v", err)
	}

	return dsn, nil
}
