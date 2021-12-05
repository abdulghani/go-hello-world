package main

import (
	"hello_world/src/router"
	"hello_world/src/utils"

	_ "github.com/99designs/gqlgen/cmd"
)

func main() {
	utils.LoadEnv()
	utils.ConnectAWS()
	r := router.InitRouter()

	r.Run(":" + utils.GetPort())
}
