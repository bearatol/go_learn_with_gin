package handlers

import (
	"net/http"
	"strconv"

	"github.com/bearatol/go_learn_with_gin/cmd/functions"
	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	data := gin.H{
		"title": "Home Page",
		"h1":    "Home Page h1",
		"year":  strconv.Itoa(functions.GetYear()),
	}
	//data = append(data, app.GlobData...)
	c.HTML(http.StatusOK, "index.html", data)

	// c.String(http.StatusOK, "<h1>Hello %d %d</h1>", 111, 222)
}
