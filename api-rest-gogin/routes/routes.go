package routes

import (
	ctrl "api-rest-gogin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", ctrl.ExibeTodosAlunos)
	r.POST("/alunos", ctrl.CriaNovoAluno)
	r.GET("/alunos/:id", ctrl.BuscaAlunoID)
	r.GET("/alunos/cpf/:cpf", ctrl.BuscaAlunoCPF)
	r.DELETE("/alunos/:id", ctrl.DeletaAluno)
	r.PATCH("/alunos/:id", ctrl.EditaAluno)
	r.GET("/:nome", ctrl.Saudacao)
	r.Run()
}
