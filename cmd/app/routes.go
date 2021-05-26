package app

import (
	"github.com/bearatol/go_learn_with_gin/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func RoutesInit() (router *gin.Engine, err error) {
	// return nil, errors.New("error server!!!")
	r := gin.Default()

	r.GET("/", handlers.MainPage)

	r.GET("/ping/", handlers.PingPage)

	return r, nil
}
