package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

type Server struct {
	Port string
}

type hfunc func(http.ResponseWriter, *http.Request) error

type Error struct {
	Error string
}

func (s *Server) Twilio(w http.ResponseWriter, r *http.Request) error {
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

func HandleFunc(f hfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			Write(w, http.StatusBadRequest, Error{Error: err.Error()})
		}
	}
}

func Write(w http.ResponseWriter, status int, msg any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("V: %s", msg)
	return json.NewEncoder(w).Encode(msg)
}

func NewServer(port string) *Server {
	s := &Server{
		Port: port,
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("STARTING SERVER ON PORT: %s \n", s.Port)

	router := mux.NewRouter()

	router.HandleFunc("/", HandleFunc(s.Twilio))
	log.Fatal(http.ListenAndServe(s.Port, router))
}
