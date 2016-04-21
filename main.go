package main

import (
	"github.com/thomas-barthelemy/go-url-shortener/server"
	"github.com/thomas-barthelemy/go-url-shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/vharitonsky/iniflags"
	"fmt"
)

func main() {
	utils.Init()

	// Getting parameters
	iniflags.Parse()
	config := utils.GetConfig()

	// Create Router Engine
	router := gin.New()

	// Setup Routes
	server.InitRoutes(router)

	router.Run(fmt.Sprintf(":%d", *config.AppPort))
}
