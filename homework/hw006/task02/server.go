package task02

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type listenServer struct {
	Address string
}

// ListenServer ...
type ListenServer interface {
	ReturnMultipliedOrUppercaseMessages(w http.ResponseWriter, r *http.Request)
	Run()
}

// NewListenServer ...
func NewListenServer(address string) ListenServer {
	return listenServer{
		Address: address,
	}
}

// ReturnMultipliedOrUppercaseMessages ...
func (l listenServer) ReturnMultipliedOrUppercaseMessages(w http.ResponseWriter, r *http.Request) {
	respBody, _ := ioutil.ReadAll(r.Body)
	separatedItems := strings.Split(string(respBody),`\n`)
	for i := 0; i < len(separatedItems); i++ {
	}
	fmt.Fprintln(w, "separatedItems: ", separatedItems)

}

func (l listenServer) Run() {
	handler := http.HandlerFunc(l.ReturnMultipliedOrUppercaseMessages)
	log.Fatal(http.ListenAndServe(l.Address, handler))
}
