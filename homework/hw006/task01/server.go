package task01

import (
	"encoding/json"
	"log"
	"net/http"
)

type echoServer struct {
	Address     string
	SomePrivate string
}

// EchoServer server command interface
type EchoServer interface {
	ShowHeaders(w http.ResponseWriter, r *http.Request)
	Run()
}

// NewEchoServer create a new server on demand
func NewEchoServer(address string) EchoServer {
	return echoServer{
		Address:     address,
		SomePrivate: "localhost:5000",
	}
}

type showHeadersResponse struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestURI string      `json:"request_uri"`
	Headers    http.Header `json:"headers"`
}

// ShowHeaders route returning headers in response
func (e echoServer) ShowHeaders(w http.ResponseWriter, r *http.Request) {
	raw := showHeadersResponse{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Headers:    r.Header,
	}
	err := json.NewEncoder(w).Encode(raw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error while processing JSON. Err: %s", err)
	}
}

// Run start a server
func (e echoServer) Run() {
	handler := http.HandlerFunc(e.ShowHeaders)
	log.Fatal(http.ListenAndServe(e.Address, handler))
}
