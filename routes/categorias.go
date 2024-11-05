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

// GetCategoriaID godoc
// @Summary      Retrieve category by ID
// @Description  Obtém uma categoria de acordo com o ID
// @Tags         Categorias
// @Accept       json
// @Produce      json
// @Param        id path int true "Category ID"
// @Success      200 {object} models.Categorias
// @Router       /categorias/{id} [get]
func GetCategoriaID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categoria models.Categorias
		id := c.Param("id")

		// Realiza a consulta no banco de dados
		err := db.Where("id = ?", id).First(&categoria).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar a categoria"})
			return
		}
		c.JSON(http.StatusOK, categoria)
	}
}

// PostCategoria godoc
// @Summary      Add a new category
// @Description  Adiciona uma nova categoria
// @Tags         Categorias
// @Accept       json
// @Produce      json
// @Param        categoria body models.Categorias true "Category object"
// @Success      201 {object} models.Categorias
// @Router       /categorias [post]
func PostCategoria(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var categoria models.Categorias

		// Faz o bind do JSON para o struct Cateforias
		if err := c.ShouldBindJSON(&categoria); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// salva a nova categoria no banco de dados
		err := db.Create(&categoria).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar a categoria"})
			return
		}

		// retorna a categoria criada
		c.JSON(http.StatusCreated, gin.H{"message": "Categoria criada com sucesso", "categoria": categoria})
	}
}
