package main

import (
	"database/sql"
	"technostrelka/internal/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "internal/db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sts := `DROP TABLE IF EXISTS users;
		DROP TABLE IF EXISTS contests;
		CREATE TABLE users (login TEXT PRIMARY KEY, password TEXT);
		CREATE TABLE contests (id INT PRIMARY KEY, title TEXT, description TEXT, link TEXT);
		INSERT INTO USERS (login, password) VALUES ("admin", "root")`
	_, err = db.Exec(sts)

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Static("/static", "./internal/static")
	r.LoadHTMLGlob("./internal/static/html/*.html")

	r.GET("/", handlers.Index)

	r.GET("/login", handlers.Login)
	r.GET("/signup", handlers.Signup)

	r.POST("/login", handlers.LoginPOST(db))
	r.POST("/signup", handlers.SignupPOST(db))

	r.GET("/profile", handlers.Profile(db))

	r.GET("/admin", handlers.Admin)
	r.POST("/admin", handlers.AddContest(db))

	r.GET("/api/contest", handlers.Contests(db))

	r.Run(":8080")
}
