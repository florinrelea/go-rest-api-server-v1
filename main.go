package main

import (
	"example.com/server"
)

func main() {
	server := server.ServerApp{}

	server.Init(":9003")
	server.Run()
}
