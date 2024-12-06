package main

import (
	"fmt"
	"log"

	"github.com/eduardomarini/API_GO_DB/db"
	_ "github.com/eduardomarini/API_GO_DB/docs"
	"github.com/eduardomarini/API_GO_DB/migrations"
	"github.com/eduardomarini/API_GO_DB/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Cria uma instância do banco de dados
	database, err := db.Connect()
	if err != nil {
		// Loga o erro e encerra o programa se a conexão falhar
		log.Fatal(err)
	}

	// Faz todas as migrações
	err = migrations.RunMigrations(database)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tabela Autores migrada com sucesso!")

	r := gin.Default()

	// rota para obter autores
	r.GET("/autores", routes.GetAutores(database))
	// rota para obter um autor por ID
	r.GET("/autores/:id", routes.GetAutorID(database))
	// rota para obter livros
	r.GET("/livros", routes.GetLivros(database))
	// rota para obter um livro por ID
	r.GET("/livros/:id", routes.GetLivroID(database))
	// rota para obter usuários
	r.GET("/usuarios", routes.GetUsuarios(database))
	// rota para obter um usuário por ID
	r.GET("/usuarios/:id", routes.GetUsuarioID(database))
	// rota para obter categorias
	r.GET("/categorias", routes.GetCategorias(database))
	// rota para obter uma categoria por ID
	r.GET("/categorias/:id", routes.GetCategoriaID(database))

	// rota para adicionar categoria
	r.POST("/categorias", routes.PostCategoria(database))
	// rota para adicionar livro
	r.POST("/livros", routes.PostLivro(database))
	// rota para adicionar autor
	r.POST("/autores", routes.PostAutor(database))
	// rota para adicionar usuário
	r.POST("/usuarios", routes.PostUsuario(database))

	//remover autores por ID
	r.DELETE("/autores/:id", routes.DeleteAutorID(database))

	// Configura a rota para a documentação do swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Executa o servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatal("erro ao iniciar o servidor: ", err)
	}
}
