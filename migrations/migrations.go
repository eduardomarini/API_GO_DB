package migrations

import (
	"fmt"

	"github.com/eduardomarini/API_GO_DB/models"
	"gorm.io/gorm"
)

// criar a tabela Autores
func MigrationsAutores(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Autores{})
	if err != nil {
		return fmt.Errorf("erro ao migrar a tabela Autores: %v", err)
	}
	fmt.Println("Tabela Autores migrada com sucesso!")
	return nil
}

// criar a tabela Categorias
func MigrationCategorias(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Categorias{})
	if err != nil {
		return fmt.Errorf("erro ao migrar a tabela Categorias: %v", err)
	}
	fmt.Println("Tabela Categorias migrada com sucesso!")
	return nil
}

// criar a tabela Emprestimos
func MigrationEmprestimos(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Emprestimos{})
	if err != nil {
		return fmt.Errorf("erro ao migrar a tabela Empréstimos: %v", err)
	}
	fmt.Println("Tabela Empréstimos migrada com sucesso!")
	return nil
}

func MigrationLivros(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Livros{})
	if err != nil {
		return fmt.Errorf("erro ao migrar a tabela Livros: %v", err)
	}
	fmt.Println("Tabela Livros migrada com sucesso!")
	return nil
}

func MigrationUsuarios(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Usuarios{})
	if err != nil {
		return fmt.Errorf("erro ao migrar a tabela Usuarios: %v", err)
	}
	fmt.Println("Tabela Usuarios migrada com sucesso!")
	return nil
}

// função para executar todas as migrações
func RunMigrations(db *gorm.DB) error {
	if err := MigrationsAutores(db); err != nil {
		return err
	}
	if err := MigrationCategorias(db); err != nil {
		return err
	}
	if err := MigrationEmprestimos(db); err != nil {
		return err
	}
	if err := MigrationLivros(db); err != nil {
		return err
	}
	if err := MigrationUsuarios(db); err != nil {
		return err
	}
	return nil
}
