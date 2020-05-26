package main

import (
	"chatserver/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":8080")

	// start the server
	s.Start()
}
