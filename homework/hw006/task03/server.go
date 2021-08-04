package task03

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	indexPath = "./static/index.html"
)

var (
	portToServe           = os.Getenv("HW006SERVICEADDRESS")
	addressOfTokenService = os.Getenv("SAVETOKENSERVICEADDR")
	portOfTokenService    = os.Getenv("SAVETOKENSERVICEPORT")
)

// WebServer ...
type WebServer struct {
	PortToServe     string
	TokenServiceURL string
	Output          func(w io.Writer, a ...interface{}) (n int, err error)
	Log             func(v ...interface{})
}

// Token ...
type Token struct {
	Name      string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

func (s *WebServer) solutionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGetMSG(w, r)
	case "POST":
		s.handlePostMSG(w, r)
	}
}

func (s *WebServer) handlePostMSG(w http.ResponseWriter, r *http.Request) {
	s.Log("POST from", r.RemoteAddr)
	var (
		deferTokenSaving = 0
		name             = r.FormValue("name")
		adr              = r.FormValue("address")
		token            = Token{
			Name:      name + ":" + adr,
			CreatedAt: time.Now(),
			ExpiredAt: time.Now().AddDate(0, 0, 10),
		}
	)
	defer func() {
		if deferTokenSaving == 1 {
			r, err := s.saveTokenToRemote(token)
			if err != nil {
				s.Log("Shit happens:", err)
				return
			}
			s.Log("Request was send:", r)
			return
		}
		s.Log("Nothing was send: no token was generated")
	}()

	err := r.ParseForm()
	if err != nil {
		s.Output(w, "ParseForm() err: %v", err)
		return
	}
	cookieToken := http.Cookie{
		Name:    "token",
		Value:   token.Name,
		Expires: token.ExpiredAt,
	}

	if name != "" && adr != "" {
		http.SetCookie(w, &cookieToken)
		s.Log("http.SetCookie:", &cookieToken)
		deferTokenSaving = 1
	}
	s.Log("redirected to /")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *WebServer) handleGetMSG(w http.ResponseWriter, r *http.Request) {
	s.Log("GET from", r.RemoteAddr)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, indexPath)
}

func (s *WebServer) saveTokenToRemote(token Token) (res *http.Response, err error) {
	data, err := json.Marshal(token)
	if err != nil {
		s.Log("json.Marshal(token) error")
		return
	}
	res, err = http.Post(
		s.TokenServiceURL,
		"application/json", bytes.NewBuffer(data))
	if err != nil {
		s.Log("post json error")
		return
	}
	return
}

func genTokenServiceURL() (u string) {
	var URL = &url.URL{}
	URL.Scheme = "http"
	URL.Host = addressOfTokenService + ":" + portOfTokenService
	URL.Path = "save-token"
	u = URL.String()
	return
}

// NewWebServer constructor for instance
func NewWebServer() *WebServer {
	return &WebServer{
		PortToServe:     ":" + portToServe,
		TokenServiceURL: genTokenServiceURL(),
		Output:          fmt.Fprintln,
		Log:             log.Println,
	}
}

// Run start a new server instance
func (s WebServer) Run() {
	fmt.Println("Using localhost:", s.PortToServe)
	r := mux.NewRouter()
	r.HandleFunc("/", s.solutionHandler).Methods("GET", "POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(s.PortToServe, nil))
}
