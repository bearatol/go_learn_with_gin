package handlers

import (
	"net/http"
	"net/url"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func AuthPagePost(c *gin.Context) {
	login := c.PostForm("LOGIN")
	password := c.PostForm("PASSWORD")
	agreement := c.PostForm("AGREEMENT")

	q := url.Values{}

	if len(login) < 1 || len(password) < 1 || len(agreement) < 1 {
		log.Info("auth form is empty")
		q.Set("isParams", "N")
		location := url.URL{Path: "/auth/", RawQuery: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
		c.Abort()
		return
	}

	users := models.GetAllUsers()
	if users == nil {
		log.Warn("user list is empty")
		c.Redirect(http.StatusMovedPermanently, "/auth/")
		c.Abort()
		return
	}
	var authLoginCreate bool
	for _, user := range users {
		if login == user.Login && password == user.Password {
			c.SetCookie("authLogin", user.Login, 0, "/", "localhost", false, true)
			authLoginCreate = true
			break
		}
	}

	if authLoginCreate {
		q.Set("isAuth", "Y")
		location := url.URL{Path: "/auth/", RawQuery: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
		c.Abort()
		return
	}

	log.Info("cookie 'authLogin' not created")
	c.Redirect(http.StatusMovedPermanently, "/auth/")
	c.Abort()
	return
}
