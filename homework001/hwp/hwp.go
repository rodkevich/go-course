package hwp

import (
	"bytes"
	"fmt"
	"github.com/kyokomi/emoji"
)

const prefix =  "Hello, world "

// 	DisplayHelloWorldWithEmoji prints `Hello, world =)`
func DisplayHelloWorldWithEmoji(s string) {
	var b bytes.Buffer
	b.WriteString(prefix)
	b.WriteString(s)
	var lineWithEmoji = emoji.Sprint(b.String())
	fmt.Println(lineWithEmoji)
	return
}
