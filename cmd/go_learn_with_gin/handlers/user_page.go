package handlers

import (
	"net/http"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/domain"
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/gin-gonic/gin"
)

func UserPage(c *gin.Context) {
	name := c.Params.ByName("name")
	data := functions.Addmap(domain.GlobalGinParams(c), gin.H{
		"title": "User page",
		"h1":    "User page",
		"user":  name,
	})
	if data["authLogin"] != name {
		Page404(c)
		return
	}

	c.HTML(http.StatusOK, "user.html", data)
}

// func UserPage2(c *gin.Context) {
// 	name := c.Param("name")
// 	action := c.Param("action")
// 	message := name + " is " + action
// 	c.String(http.StatusOK, message)
// }
