package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:nome", controller.Saudacao)
	r.GET("/alunos", controller.ExibeTodosAlunos)
	r.GET("/alunos/:id", controller.BuscaAlunoPorId)
	r.POST("/alunos", controller.CriaNovoAluno)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controller.BuscaAlunoPorCpf)
	r.GET("/alunos/", controller.BuscaAlunoPorCpf)
	r.Run()
}
