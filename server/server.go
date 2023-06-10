package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func (s *Server) Send(w http.ResponseWriter, r *http.Request) error {
	msg := NewMessage("Message")
	return Write(w, http.StatusAccepted, msg)
}
func (s *Server) Start() {
	fmt.Printf("STARTING SERVER ON PORT: %s", s.Port)

	router := mux.NewRouter()

	router.HandleFunc("/", HandleFunc(s.Send))
	log.Fatal(http.ListenAndServe(s.Port, router))
}
