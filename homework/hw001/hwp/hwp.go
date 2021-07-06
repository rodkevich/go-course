package hwp

import (
	"bytes"
	"github.com/kyokomi/emoji"
	"log"
)

const prefix = "Hello, world "

// 	AddEmoji prints `Hello, world ` with an added emoji symbol
func AddEmoji(symbol string, writer func(a ...interface{}) (n int, err error)) int {
	var b bytes.Buffer // сложение литералов через байт-буфер кажись самое оптимальное
	b.WriteString(prefix)
	b.WriteString(symbol)
	var lineWithEmoji = emoji.Sprint(b.String())
	bytesWriten, err := writer(lineWithEmoji)
	if err != nil {
		log.Fatal("Oh no, AddEmoji is broken")
	}
	return bytesWriten
}
