package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{})
}
