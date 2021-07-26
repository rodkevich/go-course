package task02

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type clientCLI struct {
	Address string
	client  http.Client
}

// NewClient ...
func NewClient(address string) Client {
	return clientCLI{
		Address: address,
		client: http.Client{
			Timeout:
			time.Second * 3,
		},
	}
}

type Client interface {
	CallServer([]byte)
	Start()
}

// CallServer ...
func (c clientCLI) CallServer(lines []byte) {
	url := "http://" + c.Address
	body := bytes.NewBuffer(lines)
	req, _ := http.NewRequest(http.MethodPost, url, body)
	resp, _ := c.client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("Response form Server: %#v\n", string(respBody))
}

// Start ...
func (c clientCLI) Start() {
	fmt.Print("Enter what u wanna process: \n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(0)
		}
		c.CallServer(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading os.Stdin:", err)
	}
}
