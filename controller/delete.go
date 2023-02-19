package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/model"
)

func DeletaAluno(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno removido com sucesso"})
}
