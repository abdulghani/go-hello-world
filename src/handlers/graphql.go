package handlers

import (
	"hello_world/src/graph"
	"hello_world/src/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func GraphQL(c *gin.Context) {
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	handler.ServeHTTP(c.Writer, c.Request)
}
