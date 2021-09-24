package router

import (
	"geocoder-v2/handler"
	"geocoder-v2/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//router.GET("/user/:name",handler.Upload)
	router.Use(middleware.Cors())
	router.POST("/upload", handler.Upload)
	return router
}
