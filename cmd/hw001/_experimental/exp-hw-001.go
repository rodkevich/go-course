package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
	hwp "github.com/rodkevich/go-course/homework/hw001/helloWorldPrinter"
)

func main() {
	// using bytes.Buffer
	var w = fmt.Println
	hwp.AddEmoji(`:relaxed:`, w)

	// using + operator
	var p = emoji.Println
	hwp.AddEmoji2(`:relaxed:`, p)

	// using strings.Builder
	hwp.AddEmoji3(`:relaxed:`, p)
}
