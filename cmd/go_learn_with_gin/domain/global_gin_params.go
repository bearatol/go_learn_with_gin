package domain

import (
	"strconv"

	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/functions"
	"github.com/gin-gonic/gin"
)

func GlobalGinParams(c *gin.Context) map[string]interface{} {
	authLogin, _ := c.Cookie("authLogin")
	return map[string]interface{}{
		"year":      strconv.Itoa(functions.GetYear()),
		"authLogin": authLogin,
	}
}
