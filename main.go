package main

import "geocoder-v2/router"

func main() {
	engine := router.SetupRouter()
	_ = engine.Run()
}
