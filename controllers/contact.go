package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Contact(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Contact page",
	})
}
