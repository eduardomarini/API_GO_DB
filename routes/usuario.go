package routes

import (
	"net/http"

	"github.com/eduardomarini/API_GO_DB/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsuarios godoc
// @Summary 	Retrieve all users
// @Description 	Obtém uma lista de todos os usuários
// @Tags 		Usuários
// @Accept 		json
// @Produce 		json
// @Success 	200 {array} models.Usuarios
// @Router 		/usuarios [get]
func GetUsuarios(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuarios []models.Usuarios

		// Realiza a consulta no banco de dados
		if err := db.Find(&usuarios).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar usuários"})
			return
		}
		// Retorna os usuálrios como json
		c.JSON(http.StatusOK, usuarios)

	}
}

// GetUsuarioID godoc
// @Summary 	Retrieve user by ID
// @Description 	Obtém um usuário de acordo com o ID
// @Tags 		Usuários
// @Accept 		json
// @Produce 		json
// @Param 		id path int true "User ID"
// @Success 	200 {object} models.Usuarios
// @Router 		/usuarios/{id} [get]
func GetUsuarioID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuario models.Usuarios
		id := c.Param("id")

		// Realiza a consulta no banco de dados
		if err := db.Where("id = ?", id).First(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar o usuário"})
			return
		}
		c.JSON(http.StatusOK, usuario)
	}
}
