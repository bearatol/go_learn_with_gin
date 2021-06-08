package handlers

import (
	"net/http"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/domain"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/models"
	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	data := functions.Addmap(domain.GlobalGinParams(c), gin.H{
		"title":    "Home Page",
		"h1":       "Home Page h1",
		"articles": models.GetAllArticles(),
	})
	c.HTML(http.StatusOK, "index.html", data)
}
