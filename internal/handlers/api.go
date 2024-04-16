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

		_, err = db.Exec("INSERT INTO contests (id, title, description, link, img) VALUES ($1, $2, $3, $4, $5)", lastId, jsn["title"], jsn["description"], jsn["link"], jsn["img"])
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

func Contests(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("select * from contests")
		if err != nil {
			panic(err)
		}

		arr := []any{}

		for rows.Next() {
			contest := struct {
				id          int
				title       string
				description string
				link        string
				img         string
			}{}
			err := rows.Scan(&contest.id, &contest.title, &contest.description, &contest.link, &contest.img)
			if err != nil {
				fmt.Println(err)
				continue
			}
			arr = append(arr, map[string]any{
				"id":          contest.id,
				"title":       contest.title,
				"description": contest.description,
				"link":        contest.link,
				"img":         contest.img,
			})
		}
		fmt.Println(arr)

		c.JSON(http.StatusOK, arr)
	}
}

func UserInfo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bts, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		jsn := map[string]string{}
		json.Unmarshal(bts, &jsn)

		fmt.Println(jsn)

		user := struct {
			login      string
			codeforces string
			acmp       string
			yandex     string
		}{
			login: jsn["login"],
		}

		row := db.QueryRow("select codeforces, acmp, yandex from users where login = $1", jsn["login"])
		err = row.Scan(&user.codeforces, &user.acmp, &user.yandex)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"login": user.login,
			"codeforces": user.codeforces,
			"acmp": user.acmp,
			"yandex": user.yandex,
		})
	}
}

func EditUserAccounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bts, err := io.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}

		jsn := map[string]string{}
		json.Unmarshal(bts, &jsn)

		fmt.Println(jsn)

		_, err = db.Exec("UPDATE users SET codeforces = $1, acmp = $2, yandex = $3 where login = $4", jsn["codeforces"], jsn["acmp"], jsn["yandex"], jsn["login"])
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{})
	}
}
