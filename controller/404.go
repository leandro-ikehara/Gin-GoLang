package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
