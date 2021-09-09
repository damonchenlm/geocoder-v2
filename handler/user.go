package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSave(context *gin.Context) {
	name := context.Param("name")
	context.String(http.StatusOK, name + "已保存")
}

