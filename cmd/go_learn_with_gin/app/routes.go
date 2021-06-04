package app

import (
	"github.com/bearatol/go_learn_with_gin/cmd/go_learn_with_gin/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesInit() (router *gin.Engine, err error) {
	// return nil, errors.New("error server!!!")
	r := gin.Default()

	r.Static("/public", "cmd/go_learn_with_gin/public")
	// r.StaticFS("/public", http.Dir("cmd/public"))
	r.StaticFile("/favicon.ico", "cmd/go_learn_with_gin/public/images/favicon.ico")

	r.LoadHTMLGlob("cmd/go_learn_with_gin/templates/*")

	r.NoRoute(handlers.Page404)

	r.GET("/", handlers.MainPage)
	r.GET("/auth/", handlers.AuthPage)
	r.POST("/auth/", handlers.AuthPagePost)

	r.GET("/user/:name/", handlers.UserPage)

	// r.GET("/user/:name/*action", handlers.UserPage2)

	r.GET("/ping/", handlers.PingPage)

	return r, nil
}
