package task01

import (
	"encoding/json"
	"log"
	"net/http"
)

type EchoServer struct {
}

type ShowHeadersResponse struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestUri string      `json:"request_uri"`
	Headers    http.Header `json:"headers"`
}

func (e *EchoServer) ShowHeaders(w http.ResponseWriter, r *http.Request) {
	raw := ShowHeadersResponse{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestUri: r.RequestURI,
		Headers:    r.Header,
	}
	err := json.NewEncoder(w).Encode(raw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error on JSON marshal. Err: %s", err)
	}
}

func (e *EchoServer) Run() {
	handler := http.HandlerFunc(e.ShowHeaders)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
