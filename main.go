package main

import (
	"fmt"
	"os"

	"github.com/kobi824/laughing-octo-funicular/server"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	s := server.NewServer(":3000")
	s.Start()
	//id := os.Getenv("ACCOUNT_SID")
	//tok := os.Getenv("TOKEN")
	client := twilio.NewRestClient()

	p := &api.CreateMessageParams{}
	p.SetBody("This is my message")
	p.SetFrom(os.Getenv("NUM2"))
	p.SetTo(os.Getenv("NUM"))

	r, err := client.Api.CreateMessage(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	if r.Sid != nil {
		fmt.Println(*r.Sid)
	} else {
		fmt.Println(r.Sid)
	}
}
