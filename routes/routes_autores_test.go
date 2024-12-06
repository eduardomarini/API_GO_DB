package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardomarini/API_GO_DB/db"
	"github.com/eduardomarini/API_GO_DB/models"
	"github.com/eduardomarini/API_GO_DB/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAutores(t *testing.T) {
	testeDB, err := SetupTestDB()
	if err != nil {
		t.Fatalf("erro ao configurar o banco de dados de teste: %v", err)
	}
	defer db.Close(testeDB)

	// Limpa o banco de dados antes de popular os dados de teste
	if err := utils.ResetTestAutoresDB(testeDB); err != nil {
		t.Fatalf("erro ao limpar o banco de dados de teste: %v", err)
	}

	testAutor := []models.Autores{
		{
			IDAutor:         1,
			Nome:            "Joana",
			Nacionalidade:   "Brasileira",
			Data_nascimento: utils.ParseDate("1990-01-01"),
		},
		{
			IDAutor:         2,
			Nome:            "Maria",
			Nacionalidade:   "Brasileira",
			Data_nascimento: utils.ParseDate("1995-01-01"),
		},
	}
	for _, autor := range testAutor {
		testeDB.Create(&autor)
	}
	// Configurar o router do Gin e registra a rota para GetAutores
	r := gin.Default()
	r.GET("/autores", GetAutores(testeDB))

	// Cria uma requisição HTTP de teste
	req, _ := http.NewRequest("GET", "/autores", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// verifica o corpo da resposta
	expectedJSON := `[{	
		"ID_autor": 1,
		"Nome": "Joana",
		"Nacionalidade": "Brasileira",
		"Data_nascimento": "1990-01-01"
		}, {
		"ID_autor": 2,
		"Nome": "Maria",
		"Nacionalidade": "Brasileira",
		"Data_nascimento": "1995-01-01"
	}]`

	// compara o JSON esperado com o JSON retornado
	assert.Equal(t, expectedJSON, w.Body.String())
}
