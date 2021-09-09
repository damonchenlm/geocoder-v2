package router

import (
	"geocoder-v2/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/user/:name",handler.UserSave)
	return router
}
