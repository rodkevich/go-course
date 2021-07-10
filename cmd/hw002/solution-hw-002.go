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
	"flag"
	"fmt"
	"os"

	"github.com/rodkevich/go-course/homework/hw002/fibo"
)

var (
	err error
)

//handleError
// TODO: разобраться когда нужно/не нужно выходить с 1
func handleError(exitCode int, err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(exitCode)
	}
}

func main() {
	defer fmt.Println("\n<< Main program exited")
	experimentalPart := flag.Bool("e", false, "Run with experiments")
	flag.Parse()
	var n int
	fmt.Print("Enter an integer between 1 - 93: \n")
	_, err = fmt.Scanln(&n) // get input value into n var
	handleError(1, err)     // make app exit with 1 exit code if err not nil
	//Solution part
	handleError(0, fibo.Solution(n))

	/* Solution(n) output for n = 25:
	Enter an integer between 1 - 93:
	25
	Solution part started >>

	Counted fresh sequence part
	Result:
	[0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368]
	<< Solution part ended

	<< Main program exited
	*/

	//////////////////  Experimental   /////////////////////
	//////////////////      part       /////////////////////
	//☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣
	if *experimentalPart {
		fmt.Println("Experimental part started >>")
		defer fmt.Println("<< Experimental part ended")
		cache := fibo.NewCache()
		handleError(0, fibo.SequenceWithCaching(int64(n), &cache))
		handleError(0, fibo.SequenceWithCaching(7, &cache))
		handleError(0, fibo.SequenceWithCaching(4, &cache))
		handleError(0, fibo.SequenceWithCaching(20, &cache))
		handleError(0, fibo.SequenceWithCaching(17, &cache))

		/* Solution(n) +  OptionalExperiments(n) output for n = 5:
		Enter an integer between 1 - 93:
		11
		Solution part started >>

		Counted fresh sequence part
		Result:
		[0 1 1 2 3 5 8 13 21 34 55]
		<< Solution part ended
		Experimental part started >>

		Counted fresh sequence part
		Result:
		[0 1 1 2 3 5 8 13 21 34 55]

		Using cache
		Result:
		[0 1 1 2 3 5 8]

		Using cache
		Result:
		[0 1 1 2]

		Counted fresh sequence part
		Result:
		[0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181]

		Using cache
		Result:
		[0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987]
		<< Experimental part ended

		<< Main program exited
		*/
	}
}
