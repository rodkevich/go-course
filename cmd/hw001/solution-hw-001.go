/*
Lesson_01:
Create new application
Use emoji package "github.com/kyokomi/emoji" to say `Hello, world :relaxed:` with emoji
Run it
*/

package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

// main prints `Hello, world :relaxed:` with emoji
func main() {

	fmt.Println(emoji.Sprint(`Hello, world :relaxed:`))
}

// Result:
// Hello, world ☺️
