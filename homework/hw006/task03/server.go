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

// NewWebServer constructor for instance
func NewWebServer() *webServer {
	return &webServer{
		Address: address,
		output:  fmt.Fprintln,
	}
}

// Run start a new server instance
func (s webServer) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/", s.solutionHandler).Methods("GET", "POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(s.Address, nil))
	// log.Fatal(http.ListenAndServe(":5050", http.FileServer(http.Dir("./static"))))
}

func (s webServer) solutionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		s.processGetMSG(w, r)
	case "POST":
		s.processPostMSG(w, r)
	default:
		s.output(w, "Use GET or POST")
	}
}

func (s webServer) processPostMSG(w http.ResponseWriter, r *http.Request) {
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

func (s webServer) processGetMSG(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, indexPath)
}
