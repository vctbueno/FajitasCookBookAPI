package main

import (
	"log"

	"github.com/webgoprojects/go-gin-template/middleware"
	//"github.com/webgoprojects/go-gin-template/models"
	"github.com/webgoprojects/go-gin-template/routers"
)

func main() {
	r := routers.SetupRouter()

	//	models.ConnectDatabase()

	// Set trusted proxies
	_ = r.SetTrustedProxies([]string{"192.168.0.1"})

	r.LoadHTMLGlob("templates/*")

	// Middleware
	r.Use(middleware.Logger())

	// Serve static files
	r.Static("/static", "./static")

	// Start the server
	log.Println("Starting server on :8080")
	r.Run(":8080")
}
