package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
	helloWorldPrinter "github.com/rodkevich/go-course/homework/hw001/hwp"
)

func main() {
	// using bytes.Buffer
	var w = fmt.Println
	helloWorldPrinter.AddEmoji(`:relaxed:`, w)

	// using + operator
	var p = emoji.Println
	helloWorldPrinter.AddEmoji2(`:relaxed:`, p)

	// using strings.Builder
	helloWorldPrinter.AddEmoji3(`:relaxed:`, p)
}
