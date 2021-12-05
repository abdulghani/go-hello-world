package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func getToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")

	if strings.TrimSpace(token) == "" {
		token = c.GetHeader("authorization")
	}

	return token
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := VerifyToken(getToken((c)))

		Inspect("AUTH TOKEN", token, err)

		c.Next()
	}
}
