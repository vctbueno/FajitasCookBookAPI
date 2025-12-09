package controllers

import (
        "net/http"

        "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
                "Title": "Home Page",
        })
}

func About(c *gin.Context) {
        c.HTML(http.StatusOK, "about.html", gin.H{
                "Title": "About Page",
        })
}