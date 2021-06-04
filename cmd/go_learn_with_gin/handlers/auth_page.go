package handlers

import (
	"net/http"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/domain"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func AuthPage(c *gin.Context) {
	if exit := c.Query("exit"); exit == "Y" {
		c.SetCookie("authLogin", "", -1, "/", "localhost", false, true)
		log.Info("cookie 'authLogin' was deleted")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
		return
	}

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
