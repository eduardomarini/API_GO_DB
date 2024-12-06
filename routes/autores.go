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

		fmt.Println("autores", autores)
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
			//fmt.Println("erro ao recuperar o autor pelo ID", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar o autor"})
			return
		}

		c.JSON(http.StatusOK, autor)
		//fmt.Print("autor", autor)

		//fmt.Println("autor", autor)
	}
}

// PostAutor godoc
// @Summary      Add a new author
// @Description  Adiciona um novo autor
// @Tags         Autores
// @Accept       json
// @Produce      json
// @Param        autor body models.Autores true "Author object"
// @Success      201 {object} models.Autores
// @Router       /autores [post]
func PostAutor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var autor models.Autores

		// Faz o bind JSON para o struct Autores
		if err := c.ShouldBindJSON(&autor); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// salva o novo autor no banco de dados
		if err := db.Create(&autor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o autor"})
			return
		}

		// retorna o autor criado
		c.JSON(http.StatusCreated, gin.H{"message": "Autor criado com sucesso", "autor": autor})

		//fmt.Println("autorPOST", autor)
	}
}

// PutAutorID godoc
// @Summary      Remove author by ID
// @Description  Remove um autor de acordo com o ID
// @Tags         Autores
// @Accept       json
// @Produce      json
// @Param        id path int true "Author ID"
// @Success      200 {object} models.Autores
// @Router       /autores/{id} [remove]
func DeleteAutorID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var autor models.Autores
		id := c.Param("id")

		// Remove o autor através do ID
		err := db.Where("id = ?", id).First(&autor).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar a categoria"})
			return
		}
		c.JSON(http.StatusOK, autor)
	}
}
