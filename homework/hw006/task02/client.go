package task02

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type clientCLI struct {
	Address string
	client  http.Client
	printer func(w io.Writer, a ...interface{}) (n int, err error)
}

// Client to work from CLI
type Client interface {
	makeCallToServer([]byte) error
	Start()
}

// NewClient constructor function
func NewClient(address string) Client {
	return clientCLI{
		Address: address,
		client: http.Client{
			Timeout: time.Second * 3,
		},
		printer: fmt.Fprintln,
	}
}

// makeCallToServer function connects and requests remote
func (c clientCLI) makeCallToServer(lines []byte) (err error) {

	url := "http://" + c.Address
	body := bytes.NewBuffer(lines)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return
	}
	rtn, err := c.client.Do(req)
	if err != nil {
		return
	}
	rtnBodyBytes, err := ioutil.ReadAll(rtn.Body)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(rtn.Body)
	_, err = c.printer(os.Stdout, string(rtnBodyBytes))
	if err != nil {
		return
	}
	return
}

// Start the instance of client
func (c clientCLI) Start() {
	fmt.Print("Enter what u wanna process: \n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(0)
		}
		c.makeCallToServer(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		c.printer(os.Stderr, "reading os.Stdin:", err)
	}
}
