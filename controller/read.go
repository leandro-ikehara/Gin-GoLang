package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/model"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []model.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno model.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&model.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
