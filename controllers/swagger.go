package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Swagger(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "API Documentation",
	})
}
