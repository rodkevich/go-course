package task01

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type echoServer struct {
	Address     string
	SomePrivate string
}

// EchoServer server command interface
type EchoServer interface {
	ReturnHeaders(w http.ResponseWriter, r *http.Request)
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

// ReturnHeaders returns headers in response
func (e echoServer) ReturnHeaders(w http.ResponseWriter, r *http.Request) {
	raw := showHeadersResponse{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Headers:    r.Header,
	}
	err := json.NewEncoder(w).Encode(raw)
	if err != nil {
		fmt.Fprintln(w, "Error while processing JSON. Err: ", err)
	}
}

// Run start a server
func (e echoServer) Run() {
	handler := http.HandlerFunc(e.ReturnHeaders)
	log.Fatal(http.ListenAndServe(e.Address, handler))
}
