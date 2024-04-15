package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func LoginPOST(db *sql.DB) gin.HandlerFunc {
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
			fmt.Println("no such user")

			if !userAndPasswordCheck(db, res["login"], res["password"]) {
				fmt.Println("password incorrect")
				c.JSON(http.StatusTeapot, gin.H{})
			} else {
				c.JSON(http.StatusOK, gin.H{})
			}
		} else {
			c.JSON(http.StatusTeapot, gin.H{})
		}
	}
}

func userAndPasswordCheck(db *sql.DB, login, target string) bool {
	row := db.QueryRow("select password from users where login = $1", login)
	password := ""
	err := row.Scan(&password)
	if err != nil {
		panic(err)
	}
	fmt.Println(password)

	return password == target
}
