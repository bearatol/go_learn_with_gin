package handlers

import (
	"net/http"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/domain"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	data := functions.Addmap(domain.GlobalGinParams(), gin.H{
		"title": "Home Page",
		"h1":    "Home Page h1",
	})
	c.HTML(http.StatusOK, "index.html", data)

	// c.String(http.StatusOK, "<h1>Hello %d %d</h1>", 111, 222)
}