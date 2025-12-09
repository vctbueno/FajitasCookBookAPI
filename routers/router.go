package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/webgoprojects/go-gin-template/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/", controllers.Index)
	r.GET("/about", controllers.About)
	r.GET("/contact", controllers.Contact)

	return r
}
