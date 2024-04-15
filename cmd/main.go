package main

import (
	"github.com/gin-gonic/gin"
	"technostrelka/internal/handlers"
)

func main() {	
	r := gin.Default()

	r.Static("/static", "./internal/static")
	r.LoadHTMLGlob("./internal/static/html/*.html")

	r.GET("/", handlers.Index)
	
	r.GET("/login", handlers.Login)
	r.GET("/signup", handlers.Signup)

	r.POST("/login", handlers.LoginPOST)
	r.POST("/signup", handlers.SignupPOST)

	r.Run(":8080")
}
