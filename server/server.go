package server

import (
	"fmt"
	"laughing-octo-funicular/twilio"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Port string
}

type hfunc func(http.ResponseWriter, *http.Request) error

type Error struct {
	Error string
}

func HandleFunc(f hfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			twilio.Write(w, http.StatusBadRequest, Error{Error: err.Error()})
		}
	}
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

	router.HandleFunc("/", HandleFunc(twilio.Twilio))
	log.Fatal(http.ListenAndServe(s.Port, router))
}
