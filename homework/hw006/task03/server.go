package task03

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	address   = "127.0.0.1:5050"
	indexPath = "./static/index.html"
)

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
	// log.Fatal(http.ListenAndServe(":5050", http.FileServer(http.Dir("./static"))))
}

func (s webServer) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		s.ifGet(w, r)
	case "POST":
		s.ifPost(w, r)
	default:
		s.output(w, "Use GET or POST")
	}
}

func (s webServer) ifPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	adr := r.FormValue("address")
	hC := http.Cookie{
		Name:  "token",
		Value: name + ":" + adr,
	}
	http.SetCookie(w, &hC)
	http.ServeFile(w, r, indexPath)
}

func (s webServer) ifGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, indexPath)
}
