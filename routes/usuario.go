package routes

import (
	"net/http"
	"time"

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

// Postusuario godoc
// @summary Add a new user
// @Description Adiciona umo novo usuário
// @Tags Usuários
// @Accept json
// @Produce json
// @Param usuario body models.Usuarios true "User object"
// @Success 201 {object} models.Usuarios
// @Router /usuarios [post]
func PostUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var usuario models.Usuarios

		// Faz o bind do Json para o struct Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Valida o formato da data de registro no formato "YYYY-MM-DD"
		if usuario.Data_registro != "" {
			if _, err := time.Parse("2006-01-02", usuario.Data_registro); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Data de registro inválida"})
				return
			}
		}

		// salva o novo usuário no banco de dados
		if err := db.Create(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
			return
		}

		// retorna o usuário criado
		c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso", "usuario": usuario})

	}
}
