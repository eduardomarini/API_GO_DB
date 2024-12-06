package utils

import (
	"fmt"
	"time"

	"github.com/eduardomarini/API_GO_DB/models"
	"gorm.io/gorm"
)

// ParseDate converte uma string no formato "2006-01-02" para um ponteiro de time.Time
func ParseDate(date string) *time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil
	}
	return &t
}

// Limpa o banco de dados de teste após os testes
func ResetTestUsuariosDB(testDB *gorm.DB) error {
	if err := testDB.Migrator().DropTable(&models.Usuarios{}); err != nil {
		return fmt.Errorf("erro ao limpar o banco de dados de teste: %v", err)
	}

	// recria a tabela de usuários para os próximos testes
	if err := testDB.AutoMigrate(&models.Usuarios{}); err != nil {
		return fmt.Errorf("erro ao recriar a tabela de usuários: %v", err)
	}
	return nil
}

// Limpa o banco de dados de teste após os testes
func ResetTestAutoresDB(testDB *gorm.DB) error {
	if err := testDB.Migrator().DropTable(&models.Autores{}); err != nil {
		return fmt.Errorf("erro ao limpar o banco de dados de teste: %v", err)
	}

	// recria a tabela de usuários para os próximos testes
	if err := testDB.AutoMigrate(&models.Autores{}); err != nil {
		return fmt.Errorf("erro ao recriar a tabela de autores: %v", err)
	}
	return nil
}
