package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func SignupPOST(c *gin.Context) {
	bts, err := io.ReadAll(c.Request.Body)

	if err != nil {
		panic(err)
	}

	res := map[string]any{}
	json.Unmarshal(bts, &res)

	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{})
}
