package task02

import (
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
	CallServerControlled(newLine string)
	Start()
}

func (c clientCLI) CallServerControlled(s string) {
	url := "http://" + c.Address
	data := s
	body := bytes.NewBufferString(data)
	req, _ := http.NewRequest(http.MethodPost, url, body)
	resp, _ := c.client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("CallServerControlled %#v\n\n\n", string(respBody))
}

func (c clientCLI) Start() {
	var newLine string
	fmt.Print("Enter what u wanna process: \n")
	_, _ = fmt.Scanln(&newLine) // get input value into newLine var
	c.CallServerControlled(newLine)
	if newLine == "exit" {
		os.Exit(0)
	}
}
