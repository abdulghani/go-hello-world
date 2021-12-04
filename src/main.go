package main

import (
	"hello_world/src/router"
	"hello_world/src/utils"
)

func main() {
	utils.LoadEnv()
	r := router.InitRouter()

	r.Run(":" + utils.GetPort())
}
