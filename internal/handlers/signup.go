package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func SignupPOST(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		bts, err := io.ReadAll(c.Request.Body)

		if err != nil {
			panic(err)
		}

		res := map[string]string{}
		json.Unmarshal(bts, &res)

		fmt.Println(res)

		if userExists(db, res["login"]) {
			fmt.Println("there's such user")
			c.JSON(http.StatusTeapot, gin.H{})
		} else {
			fmt.Println("no such user")

			_, err := db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", res["login"], res["password"])
			if err != nil {
				panic(err)
			}
			
			c.JSON(http.StatusOK, gin.H{})
		}
	}
}

func userExists(db *sql.DB, login string) bool {
	sqlStmt := `SELECT login FROM users WHERE login = ?`
	err := db.QueryRow(sqlStmt, login).Scan(&login)
	if err != nil {
		if err != sql.ErrNoRows {
			panic(err)
		}

		return false
	}
	return true
}
