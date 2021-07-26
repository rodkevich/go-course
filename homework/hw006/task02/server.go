package task02

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type listenServer struct {
	Address string
}

// ListenServer ...
type ListenServer interface {
	ReturnMultipliedOrUppercase(w http.ResponseWriter, r *http.Request)
	Run()
}

// NewListenServer ...
func NewListenServer(address string) ListenServer {
	return listenServer{
		Address: address,
	}
}

// ReturnMultipliedOrUppercase ...
func (l listenServer) ReturnMultipliedOrUppercase(w http.ResponseWriter, r *http.Request) {
	respBody, ioErr := ioutil.ReadAll(r.Body)
	if ioErr != nil {
		fmt.Fprintln(w, "reading body from request:", ioErr)
	}
	separatedEnts := strings.Split(string(respBody), `\n`)
	for i := 0; i < len(separatedEnts); i++ {
		if _, err := strconv.Atoi(separatedEnts[i]); err == nil {
			fmt.Printf("%q looks like a number.\n", separatedEnts[i])
			var a int
			fmt.Sscanf(separatedEnts[i], "%d", &a)
			fmt.Fprintln(w, a, "to Multiplied by 2:", a*2)
			continue
		}
		fmt.Fprintln(w, separatedEnts[i], "to Uppercase:", strings.ToUpper(separatedEnts[i]))
	}
}

func (l listenServer) Run() {
	handler := http.HandlerFunc(l.ReturnMultipliedOrUppercase)
	log.Fatal(http.ListenAndServe(l.Address, handler))
}
