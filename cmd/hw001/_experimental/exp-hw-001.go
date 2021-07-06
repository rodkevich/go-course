package main

import (
	"fmt"
	helloWorldPrinter "github.com/rodkevich/go-course/homework/hw001/hwp"
)

func main() {
	var w = fmt.Println
	helloWorldPrinter.AddEmoji(`:relaxed:`, w)
}
