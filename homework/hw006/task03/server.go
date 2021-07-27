package task03

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

const address = "127.0.0.1:5050"

type webServer struct {
	Address string
	output  func(w io.Writer, a ...interface{}) (n int, err error)
}

func NewWebServer() *webServer {
	return &webServer{
		Address: address,
		output:  fmt.Fprintln,
	}
}

func (s webServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.handler).Methods("GET", "POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(s.Address, nil))
}

func (s webServer) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		s.ifGet(w, r)
	case "POST":
		s.ifPost(w, r)
	}
}

func (s webServer) ifPost(w http.ResponseWriter, r *http.Request) {
	s.output(w, "Result:", "POST ...")
}

func (s webServer) ifGet(w http.ResponseWriter, r *http.Request) {
	s.output(w, "Result:", "GET ...")
}
