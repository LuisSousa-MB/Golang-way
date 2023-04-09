package routes

import (
	ctrl "api-rest-gogin/controllers"

	docs "api-rest-gogin/docs"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swaggerCaronte.yaml",
		SpecPath:    "/swagger.yaml",
		DocsPath:    "/docs",
	}

	r := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", ctrl.ExibeTodosAlunos)
	r.POST("/alunos", ctrl.CriaNovoAluno)
	r.GET("/alunos/:id", ctrl.BuscaAlunoID)
	r.GET("/alunos/cpf/:cpf", ctrl.BuscaAlunoCPF)
	r.GET("/alunos/rg/:rg", ctrl.BuscaAlunoRg)
	r.DELETE("/alunos/:id", ctrl.DeletaAluno)
	r.PATCH("/alunos/:id", ctrl.EditaAluno)
	//r.GET("/:nome", ctrl.Saudacao)
	r.GET("/index", ctrl.ExibePagIndex)
	r.GET("/docx", ctrl.ExibeDocs)
	r.GET("/swagger", ctrl.ExibeJson)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	//r.NoRoute(ctrl.Limbo)
	r.Use(ginredoc.New(doc))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	r.Run()
}
