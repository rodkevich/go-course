package main

import (
	"github.com/rodkevich/go-course/homework/hw006/task02"
)

const address = "127.0.0.1:5000"

func main() {
	l := task02.NewListenServer(address)
	go l.Run()
	cli := task02.NewClient(address)
	for {
		cli.Start()
	}
}
