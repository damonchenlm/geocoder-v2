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
	// 上传文件
	router.POST("/upload", handler.Upload)
	// 地点解析
	router.GET("/geocode",handler.Geocode)
	return router
}
