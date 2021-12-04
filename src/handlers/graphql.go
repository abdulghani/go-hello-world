package handlers

import (
	"hello_world/src/graph"
	"hello_world/src/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
)

const mb int64 = 1 << 20

func GraphQL(c *gin.Context) {
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// add multipart capability to handler
	handler.AddTransport(transport.POST{})
	handler.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 100 * mb,
	})

	handler.ServeHTTP(c.Writer, c.Request)
}
