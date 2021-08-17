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
	log     func(v ...interface{})
}

// NewWebServer constructor for instance
func NewWebServer() *webServer {
	return &webServer{
		Address: address,
		output:  fmt.Fprintln,
		log:     log.Println,
	}
}

// Run start a new server instance
func (s webServer) Run() {
	fmt.Println("Using ", address)
	r := mux.NewRouter()
	r.HandleFunc("/", s.solutionHandler).Methods("GET", "POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(s.Address, nil))
}

func (s webServer) solutionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.processGetMSG(w, r)
	case "POST":
		s.processPostMSG(w, r)
	}
}

func (s webServer) processPostMSG(w http.ResponseWriter, r *http.Request) {
	s.log("POST from", r.RemoteAddr)
	err := r.ParseForm()
	if err != nil {
		s.output(w, "ParseForm() err: %v", err)
		return
	}
	name := r.FormValue("name")
	adr := r.FormValue("address")
	token := http.Cookie{
		Name:  "token",
		Value: name + ":" + adr,
	}
	if name != "" && adr != "" {
		http.SetCookie(w, &token)
		s.log("http.SetCookie:", &token)
	}
	s.log("redirected to /")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s webServer) processGetMSG(w http.ResponseWriter, r *http.Request) {
	s.log("GET from", r.RemoteAddr)
	http.ServeFile(w, r, indexPath)
}
