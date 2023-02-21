package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandro-ikehara/Gin-GoLang/database"
	"github.com/leandro-ikehara/Gin-GoLang/model"
)

func ExibePaginaIndex(c *gin.Context) {
	var alunos []model.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}
