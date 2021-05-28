package main

import (
	"fmt"

	"github.com/bearatol/go_learn_with_gin/cmd/app"
	"github.com/labstack/gommon/log"
)

func main1() {
	router, err := app.RoutesInit()
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println("Run server http://localhost:8833")
	router.Run(":8833")
}
