package router

import (
	"hello_world/src/utils"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(router *gin.Engine) {
	router.Use(utils.AuthMiddleware())
}
