package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func generator(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "generator.html", gin.H{})
	}
}
