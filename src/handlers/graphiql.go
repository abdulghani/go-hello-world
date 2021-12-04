package handlers

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func GraphiQL(c *gin.Context) {
	handler := playground.Handler("GRAPHQL PLAYGROUND", "/graphql")
	handler.ServeHTTP(c.Writer, c.Request)
}
