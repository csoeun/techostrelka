package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var lastId = 0

func AddContest(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bts, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		jsn := map[string]string{}
		json.Unmarshal(bts, &jsn)

		fmt.Println(jsn)

		_, err = db.Exec("INSERT INTO contests (id, title, description, link) VALUES ($1, $2, $3, $4)", lastId, jsn["title"], jsn["description"], jsn["link"])
		if err != nil {
			panic(err)
		}
		lastId++

		c.JSON(http.StatusOK, gin.H{})
	}
}

func RemoveContest(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bts, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		jsn := map[string]string{}
		json.Unmarshal(bts, &jsn)

		fmt.Println(jsn)

		_, err = db.Exec("DELETE FROM contests WHERE id=$1", jsn["id"])
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}
