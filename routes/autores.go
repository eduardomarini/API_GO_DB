package routes

import (
	"fmt"
	"net/http"

	"github.com/eduardomarini/API_GO_DB/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAutores godoc
// @Summary      Retrieve all authors
// @Description  Obtém uma lista de todos os autores
// @Tags         Autores
// @Accept       json
// @Produce      json
// @Success      200 {array} models.Autores
// @Router       /autores [get]
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

		//fmt.Println("autores", autores)
	}
}

// GetAutorID godoc
// @Summary      Retrieve author by ID
// @Description  Obtém um autor de acordo com o ID
// @Tags         Autores
// @Accept       json
// @Produce      json
// @Param		id path int true "Author ID"
// @Success      200 {object} models.Autores
// @Router       /autores/{id} [get]
func GetAutorID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var autor models.Autores
		id := c.Param("id")

		// Realiza a consulta no banco de dados
		if err := db.Where("id = ?", id).First(&autor).Error; err != nil {
			fmt.Println("erro ao recuperar o autor pelo ID", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar o autor"})
			return
		}

		c.JSON(http.StatusOK, autor)
		fmt.Print("autor", autor)
	}
}
