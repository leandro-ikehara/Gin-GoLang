package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/model"
)

func CriaNovoAluno(c *gin.Context) {
	var aluno model.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
