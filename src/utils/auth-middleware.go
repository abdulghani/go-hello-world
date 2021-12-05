package utils

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("GOING THROUGH MIDDLEWARE")

		c.Next()
	}
}
