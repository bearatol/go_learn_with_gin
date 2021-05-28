package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	c.String(http.StatusOK, "<h1>Hello %d %d</h1>", 111, 222)
}
