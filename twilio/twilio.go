package twilio

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

func Twilio(w http.ResponseWriter, r *http.Request) error {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env found")
	}
	params := GetClientParams()
	client := twilio.NewRestClientWithParams(*params)
	msg := GetMessage("This is a test message")
	p := &api.CreateMessageParams{}
	p.SetBody(msg)
	p.SetFrom(os.Getenv("FROM"))
	p.SetTo(os.Getenv("TO"))

	req, err := client.Api.CreateMessage(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	if req.Sid != nil {
		fmt.Println(*req.Sid)
	} else {
		fmt.Println(req.Sid)
	}
	if http.StatusOK == 200 {
		return Write(w, http.StatusAccepted, msg)
	}
	return fmt.Errorf("there was an error")
}

func GetMessage(msg string) string {
	return msg
}

func GetClientParams() *twilio.ClientParams {
	t := &twilio.ClientParams{
		Username: os.Getenv("ACCOUNT_SID"),
		Password: os.Getenv("TOKEN"),
	}
	return t
}

func Write(w http.ResponseWriter, status int, msg any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("V: %s", msg)
	return json.NewEncoder(w).Encode(msg)
}
