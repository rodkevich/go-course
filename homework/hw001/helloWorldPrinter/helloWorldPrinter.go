package helloWorldPrinter

import (
	"bytes"
	"github.com/kyokomi/emoji"
	"log"
	"strings"
)

const prefix = "Hello, world "

//AddEmoji prints `Hello, world ` with an added emoji symbol using bytes.Buffer
func AddEmoji(symbol string, writer func(a ...interface{}) (n int, err error)) int {
	var b bytes.Buffer // сложение литералов в байт-буфер кажись самое оптимальное
	b.WriteString(prefix)
	b.WriteString(symbol)
	bytesWriten, err := writer(emoji.Sprint(b.String()))
	if err != nil {
		log.Fatal("Oh no, AddEmoji is broken")
	}
	return bytesWriten
}

//AddEmoji2 prints `Hello, world ` with an added emoji symbol using + operator
func AddEmoji2(symbol string, writer func(a ...interface{}) (n int, err error)) int {
	bytesWriten, err := writer(prefix + symbol)
	if err != nil {
		log.Fatal("Oh no, AddEmoji2 is broken")
	}
	return bytesWriten
}

//AddEmoji3 prints `Hello, world ` with an added emoji symbol using strings.Builder
func AddEmoji3(symbol string, writer func(a ...interface{}) (n int, err error)) int {
	var line strings.Builder
	line.WriteString(prefix)
	line.WriteString(symbol)
	bytesWriten, err := writer(line.String())
	if err != nil {
		log.Fatal("Oh no, AddEmoji3 is broken")
	}
	return bytesWriten
}
