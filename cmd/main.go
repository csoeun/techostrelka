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

	r.Run(":8080")
}
