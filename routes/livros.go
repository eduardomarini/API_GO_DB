package routes

import (
	"net/http"

	"github.com/eduardomarini/API_GO_DB/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetLivros godoc
// @Summary      Retrieve all books
// @Description  Obtém uma lista de todos os livros
// @Tags         Livros
// @Accept       json
// @Produce      json
// @Success      200 {array} models.Livros
// @Router       /livros [get]
func GetLivros(db *gorm.DB) gin.HandlerFunc { // Passar a instância do banco de dados como parâmetro
	return func(c *gin.Context) {
		var livros []models.Livros

		// Realiza a consulta no banco de dados
		if err := db.Find(&livros).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar os livros"})
			return
		}

		// Retorna os livros como JSON
		c.JSON(http.StatusOK, livros)
	}
}
