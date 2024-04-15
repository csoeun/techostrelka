package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddContest(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bts, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		jsn := map[string]string{}
		json.Unmarshal(bts, &jsn)

		fmt.Println(jsn)

		_, err = db.Exec("INSERT INTO contests (id, title, description, link) VALUES ($1, $2, $3, $4)", jsn["id"], jsn["title"], jsn["description"], jsn["link"])
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{})
	}
} 