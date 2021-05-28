package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserPage(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{"param": "name", "value": name})
}

func UserPage2(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}
