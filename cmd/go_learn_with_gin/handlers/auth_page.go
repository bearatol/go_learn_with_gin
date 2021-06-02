package handlers

import (
	"fmt"
	"net/http"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/domain"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/gin-gonic/gin"
)

func AuthPage(c *gin.Context) {
	authLogin, _ := c.Cookie("authLogin")
	fmt.Printf("authLogin: %v\n", authLogin)

	isParams := c.Query("isParams")
	isAuth := c.Query("isAuth")
	data := functions.Addmap(domain.GlobalGinParams(c), gin.H{
		"title":    "Auth page",
		"h1":       "Auth page",
		"isParams": isParams,
		"isAuth":   isAuth,
	})
	c.HTML(http.StatusOK, "auth.html", data)
}
