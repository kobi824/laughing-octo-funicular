package main

import (
	"fmt"
	"laughing-octo-funicular/meme"
	"laughing-octo-funicular/server"
)

func main() {
	s := server.NewServer(":3000")
	meme, _ := meme.GetRandomImage()
	fmt.Println(meme)
	s.Start()
}
