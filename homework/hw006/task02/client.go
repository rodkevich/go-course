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

// NewClient ...
func NewClient(address string) Client {
	return clientCLI{
		Address: address,
		client: http.Client{
			Timeout: time.Second * 3,
		},
		printer: fmt.Fprintln,
	}
}

// Client ...
type Client interface {
	MakeCallToServer([]byte)
	Start()
}

// MakeCallToServer ...
func (c clientCLI) MakeCallToServer(lines []byte) {

	url := "http://" + c.Address
	body := bytes.NewBuffer(lines)
	req, _ := http.NewRequest(http.MethodPost, url, body)
	res, _ := c.client.Do(req)
	rbd, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	c.printer(os.Stdout, string(rbd))
}

// Start ...
func (c clientCLI) Start() {
	fmt.Print("Enter what u wanna process: \n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(0)
		}
		c.MakeCallToServer(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		c.printer(os.Stderr, "reading os.Stdin:", err)
	}
}
