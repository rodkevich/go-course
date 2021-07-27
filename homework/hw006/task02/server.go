package task02

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/rodkevich/go-course/homework/hw006/task02/internal/helpers"
)

type listenServer struct {
	Address string
	output  func(w io.Writer, a ...interface{}) (n int, err error)
}

// ListenServer represents server
type ListenServer interface {
	procByteItemsFromStdin(w http.ResponseWriter, r *http.Request)
	Run()
}

// NewListenServer constructor
func NewListenServer(address string) ListenServer {
	return listenServer{
		Address: address,
		output:  fmt.Fprintln,
	}
}

// procByteItemsFromStdin to be used as a handler for processing requests
func (s listenServer) procByteItemsFromStdin(w http.ResponseWriter, r *http.Request) {
	rb, ioErr := ioutil.ReadAll(r.Body)
	if ioErr != nil {
		s.output(w, "reading body from request:", ioErr)
	}
	items := strings.Split(string(rb), `\n`)

	rtn := helpers.MultiplyOrUpperDepOfType(items)
	s.output(w, "Result:", rtn)
}

// Run start an instance of a server
func (s listenServer) Run() {
	handler := http.HandlerFunc(s.procByteItemsFromStdin)
	log.Fatal(http.ListenAndServe(s.Address, handler))
}
