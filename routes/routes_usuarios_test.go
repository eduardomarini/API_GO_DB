package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/eduardomarini/API_GO_DB/db"
	"github.com/eduardomarini/API_GO_DB/models"
	utils "github.com/eduardomarini/API_GO_DB/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestGetUsuarios(t *testing.T) {

	testDB, err := SetupTestDB()
	if err != nil {
		fmt.Errorf("erro ao configurar o banco de dados de teste: %v", err)
	}
	defer db.Close(testDB) // Fecha o banco de daods após os teste

	// popula o banco de dados de teste
	testUsers := []models.Usuarios{

		{
			IDUsuario:     1,
			Nome:          "Marina",
			Email:         "marinamarini@example.com",
			Telefone:      "(XX)99999-9999",
			Data_registro: utils.ParseDate("2021-09-01"),
		},
		{
			IDUsuario:     2,
			Nome:          "Eduardo",
			Email:         "eduardoribeiromarini@gmail.com",
			Telefone:      "(XX)99999-9999",
			Data_registro: utils.ParseDate("2021-09-02"),
		},
	}
	for _, user := range testUsers {
		testDB.Create(&user)
	}
	// Configurar o router do Gin e registra a rota para GetUsuarios
	r := gin.Default()
	r.GET("/usuarios", GetUsuarios(testDB))

	// Cria uma requisição HTTP GET para o endpoint "/usuarios"
	req, _ := http.NewRequest("GET", "/usuarios", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verifica o código de status da resposta
	assert.Equal(t, http.StatusOK, w.Code)

	// Verifica o corpo da resposta
	expectedJSON := `[{
			"ID_Usuario": 1,
			"Nome": "Marina",
			"Email": "marinamarini@example.com",
			"Telefone": "(XX)99999-9999",
			"Data_registro": "2021-09-01"
		}, {
			"ID_Usuario": 2,
			"Nome": "Eduardo",
			"Email": "eduardoribeiromarini@gmail.com",
			"Telefone": "(XX)99999-9999",
			"Data_registro": "2021-09-02"
	}]`

	// compara o JSON esperado com o JSON retornado
	assert.JSONEq(t, expectedJSON, w.Body.String())
}
