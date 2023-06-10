package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type Server struct {
	Port string
}

type hfunc func(http.ResponseWriter, *http.Request) error

type Message struct {
	Msg string
}

type Error struct {
	Error string
}

func NewMessage(msg string) *Message {
	m := &Message{
		Msg: msg,
	}
	return m
}

func HandleFunc(f hfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			Write(w, http.StatusBadRequest, Error{Error: err.Error()})
		}
	}
}

func Write(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewServer(port string) *Server {
	s := &Server{
		Port: port,
	}
	return s
}

func (s *Server) Twilio(w http.ResponseWriter, r *http.Request) error {
	client := twilio.NewRestClient()
	msg := GetMessage("This is a test message")
	p := &api.CreateMessageParams{}
	p.SetBody(msg)
	p.SetFrom(os.Getenv("NUM2"))
	p.SetTo(os.Getenv("NUM"))

	req, err := client.Api.CreateMessage(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	if req.Sid != nil {
		fmt.Println(*req.Sid)
	} else {
		fmt.Println(req.Sid)
	}
	return Write(w, http.StatusAccepted, msg)
}

func GetMessage(msg string) string {
	return msg
}

func (s *Server) Start() {
	fmt.Printf("STARTING SERVER ON PORT: %s", s.Port)

	router := mux.NewRouter()

	router.HandleFunc("/", HandleFunc(s.Twilio))
	log.Fatal(http.ListenAndServe(s.Port, router))
}
