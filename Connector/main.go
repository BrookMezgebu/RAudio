package main

import (
	"./indexi"
	"./server"
)

func main() {

	c := indexi.AvaliableFileSystems()
	server.StartServer(c)
}
