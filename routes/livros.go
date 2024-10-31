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

// GetLivroID godoc
// @Summary Retrieve book by ID
// @Descpription Obtém um livro de acordo com o ID
// @Tags Livros
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Livros
// @Router /livros/{id} [get]
func GetLivroID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var livro models.Livros
		id := c.Param("id")

		// Realiza a consulta no banco de dados
		if err := db.Where("id = ?", id).First(&livro).Error; err != nil {
			//fmt.Println("erro ao recuperar o livro pelo ID", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar o autor"})
			return
		}
		c.JSON(http.StatusOK, livro)
	}
}
