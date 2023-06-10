package main

import (
	"laughing-octo-funicular/server"
)

func main() {
	s := server.NewServer(":3000")
	s.Start()
}
