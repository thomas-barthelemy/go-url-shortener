package server

import (
	"github.com/gin-gonic/gin"
	"github.com/thomas-barthelemy/go-url-shortener/modules/urlShortener"
)

func InitRoutes(router *gin.Engine) {
	urlShortener.InitRoutes(router)

	api := router.Group("/api")
	{
		urlShortener.InitApiRoutes(api)
	}

}

