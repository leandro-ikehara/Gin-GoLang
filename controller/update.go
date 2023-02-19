package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/model"
)

func EditaAluno(c *gin.Context) {
	var aluno model.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)

}
