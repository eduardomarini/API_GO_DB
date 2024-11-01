package routes

import (
	"net/http"

	"github.com/eduardomarini/API_GO_DB/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCategorias godoc
// @Summary      Retrieve all categories
// @Description  Obtém uma lista de todas as categorias
// @Tags         Categorias
// '@Accept       json
// @Produce      json
// @Success      200 {array} models.Categorias
// @Router       /categorias [get]
func GetCategorias(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categorias []models.Categorias

		// Realiza a consulta no banco de dados
		if err := db.Find(&categorias).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar as categorias"})
			return
		}
		c.JSON(http.StatusOK, categorias)
	}
}
