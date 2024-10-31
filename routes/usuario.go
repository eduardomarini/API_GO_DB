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
