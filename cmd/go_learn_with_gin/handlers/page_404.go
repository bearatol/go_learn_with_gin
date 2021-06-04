package handlers

import (
	"net/http"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/domain"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func Page404(c *gin.Context) {
	data := functions.Addmap(domain.GlobalGinParams(c), gin.H{
		"title": "404 page",
		"h1":    "404 page",
	})
	c.HTML(http.StatusNotFound, "404.html", data)
	log.Warn("Failed to find route. 404 page")
}
