package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/rodkevich/go-course/homework/hw006/task01"
)

const address = "localhost:8080"

func main() {
	s := task01.NewEchoServer(address)
	go s.Run()
	// check with use of curl
	out, err := exec.Command("curl", address).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response:\n %s", out)
}

// Response:
// {"host":"localhost:8080","user_agent":"curl/7.68.0","request_uri":"/","headers":{"Accept":["*/*"],"User-Agent":["curl/7.68.0"]}}
//
// Process finished with exit code 0
