package main

import (
	"BlogProject/config"
	"BlogProject/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}
