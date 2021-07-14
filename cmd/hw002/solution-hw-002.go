package main

import (
	"fmt"
	"github.com/rodkevich/go-course/cmd/hw002/fibo"
)

func main() {
	fmt.Println("Hello from Fibonacci-app")
	defer fmt.Println("\n<< Main program exited")
	var n int64 = 93
	result, _ := fibo.SizedSequence64bit(n)
	fmt.Println(result)

}
