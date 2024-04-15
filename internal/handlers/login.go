package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginPOST(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}