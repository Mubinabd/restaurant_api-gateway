package main

import (
	// "log"
	"github.com/Mubinabd/restaurant_api-gateway/api-gateway"
	"github.com/Mubinabd/restaurant_api-gateway/api-gateway/handler"

)

func main() {

	engine := api.NewGin(handler.NewHandlerStruct())
	engine.Run(":8090")
}