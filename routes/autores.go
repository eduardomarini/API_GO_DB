package routes

import (
	"net/http"

	"github.com/eduardomarini/API_GO_DB/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAutores(db *gorm.DB) gin.HandlerFunc { // Passar a instância do banco de dados como parâmetro
	return func(c *gin.Context) {
		var autores []models.Autores

		// Realiza a consulta no banco de dados
		if err := db.Find(&autores).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar os autores"})
			return
		}

		// Retorna os autores como JSON
		c.JSON(http.StatusOK, autores)
	}
}
