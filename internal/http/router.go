package http

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	api := router.Group("/api")
	{
		api.POST("/shorten", shortenURLHandler)

	}

	return router
}
