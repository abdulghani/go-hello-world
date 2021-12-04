package handlers

import "github.com/gin-gonic/gin"

func Bark(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "woof",
	})
}
