package main

import (
	"github.com/shadowhijackers/simple-websocket/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":8081")
}
