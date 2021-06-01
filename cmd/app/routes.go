package app

import (
	"github.com/bearatol/go_learn_with_gin/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesInit() (router *gin.Engine, err error) {
	// return nil, errors.New("error server!!!")
	r := gin.Default()

	r.Static("/public", "cmd/public")
	// r.StaticFS("/public", http.Dir("cmd/public"))
	r.StaticFile("/favicon.ico", "cmd/public/images/favicon.ico")

	r.LoadHTMLGlob("cmd/templates/*")

	r.GET("/", handlers.MainPage)

	r.GET("/ping/", handlers.PingPage)

	r.GET("/user/:name", handlers.UserPage)

	r.GET("/user/:name/*action", handlers.UserPage2)

	return r, nil
}
