package main

import (
	"app/src/controller"
)

func main() {
	e := controller.Init()

	SetStaticRoute(e)
	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
