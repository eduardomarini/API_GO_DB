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
// @Description Obtém um livro de acordo com o ID
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

// PostLivro godoc
// @Summary Add a new book
// @Description Adiciona um novo livro
// @Tags Livros
// @Accept json
// @Produce json
// @Param livro body models.Livros true "Book object"
// @Success 201 {object} models.Livros
// @Router /livros [post]
func PostLivro(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var livro models.Livros

		// Faz o bind do JSON para o struct Livros
		err := c.ShouldBindJSON(&livro)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// salva o novo livro no banco de dados
		if err := db.Create(&livro).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o livro"})
			return
		}
		// Retorna o livro criado
		c.JSON(http.StatusCreated, gin.H{"message": "Livro criado com sucesso", "Livro": livro})
	}
}

// DeleteLivroID godoc
// @Summary Delete book by ID
// @Description Remove um livro de acordo com o ID
// @Tags Livros
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Livros
// @Router /livros/{id} [delete]
func DeleteLivroID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var livro models.Livros
		id := c.Param("id")

		// Consulta o livro pelo ID
		err := db.Where("id = ?", id).First(&livro).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar o livro"})
			return
		}

		// deleta o livro
		if err := db.Delete(&livro).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar o Livro"})
			return
		}
		c.JSON(http.StatusOK, livro)
	}
}
