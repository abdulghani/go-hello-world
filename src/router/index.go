package router

import (
	"hello_world/src/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	InitMiddleware(router)

	router.GET("/", handlers.Hello)
	router.GET("/ping", handlers.Handler)
	router.GET("/graphql", handlers.GraphiQL)
	router.POST("/graphql", handlers.GraphQL)

	return router
}
