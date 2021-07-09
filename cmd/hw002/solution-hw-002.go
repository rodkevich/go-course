/*
Lesson_02:
Create new application
Create a package that outputs the Fibonacci numbers to the console.
      "Printer" function should have a signature like
func name(n int)
       n int is the amount of Fibonacci numbers to output
main.go should import your package.
Print welcome message.
Use your package to print some Fibonacci numbers.
Print complete message.
Optional:
Use defer to write a complete message.
Create more than one printer function for the Fibonacci function. And use them in your main.go
*/

package main

import (
	"fmt"
	"github.com/rodkevich/go-course/homework/hw002/fibo"
)

//hwSolution  prints a sequence of fibonacci numbers up to required length
func hwSolution() {
	var n int64
	fmt.Print("Enter an integer between 1 - 93: \n")
	_, _ = fmt.Scanln(&n) // get input value into n var

	var q, err = fibo.SizedSequence64bit(n)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(q)
}

func main() {
	fmt.Println("\nProgram started ...")
	defer fmt.Println("\nProgram exited")
	hwSolution()
}

//Program started ...
//Enter an integer between 1 - 93:
//10
//Program stopped
//[0 1 1 2 3 5 8 13 21 34]
//
//Program exited
