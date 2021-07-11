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
// TODO: разобраться когда нужно/не нужно выходить с 1 а то чет наобум
func handleError(exitCode int, err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(exitCode)
	}
}

//init before running main()
//set env var defining a printer to be used
//presented: [ long, short - now FAKE UNIMPLEMENTED ]
//short can be set to true in fibo.go: var `allowedPrinters`
func init() {
	defer fmt.Println("setup: Solution-app initialized")

	err = os.Setenv("Printer", "long")
	if err != nil {
		handleError(1, err)
	}
}

func main() {

	defer fmt.Println("\n<< Main program exited")
	experimentalPart := flag.Bool("o", false, "Run with optional solution")
	flag.Parse()
	var n int
	fmt.Print("Enter an integer between 1 - 93: \n")
	_, err = fmt.Scanln(&n) // get input value into n var
	handleError(1, err)     // make app exit with 1 exit code if err not nil
	//Solution part
	fmt.Println("\n| Solution part started |")
	handleError(0, fibo.Solution(n))

	/*
	!!  Printer "short" is experimentally set to "false" in fibo.go: var allowedPrinters !!!
	So ... by default output will be ... :arrow_down:

	Solution(n) output for n = 93 [printer = "short"]:

	setup: Fibonacci-app initialized
	setup: Solution-app initialized
	Enter an integer between 1 - 93:
	93

	| Solution part started |
	! Warn: 64 bit overflow on 94 iteration
	> Required sequence part counted
	> Fibonacci-env check finished
	! Using notAllowed printer:
	! Selected printer is not implemented
	<< Solution part ended

	Abort app: implementation is not there yet. See ya!

	Process finished with exit code 0
	*/

	//////////////////  Experimental   /////////////////////
	//////////////////      part       /////////////////////
	if *experimentalPart {
		fmt.Println("\n| Experimental part started |")
		defer fmt.Println("<< Experimental part ended")
		// create a cache to be used for returns of previously counted numbers
		cache := fibo.NewCache()
		// run functions to use / not use cache
		handleError(0, fibo.OptionalSolutionWithCaching(int64(n), &cache))
		handleError(0, fibo.OptionalSolutionWithCaching(7, &cache))
		handleError(0, fibo.OptionalSolutionWithCaching(4, &cache))
		handleError(0, fibo.OptionalSolutionWithCaching(20, &cache))
		handleError(0, fibo.OptionalSolutionWithCaching(17, &cache))

		/*
		Solution(n) +  OptionalSolutionWithCaching(n) output for n = 5 [printer = "long"]:

		setup: Fibonacci-app initialized
		setup: Solution-app initialized
		Enter an integer between 1 - 93:
		5

		| Solution part started |

		> Required sequence part counted
		> Fibonacci-env check finished
		! Using long printer:
		length=5 capacity=8 Result: [0 1 1 2 3]
		<< Solution part ended

		| Experimental part started |

		> Required sequence part counted
		> Fibonacci-env check finished
		! Using long printer:
		length=5 capacity=8 Result: [0 1 1 2 3]

		> Required sequence part counted
		> Fibonacci-env check finished
		! Using long printer:
		length=7 capacity=8 Result: [0 1 1 2 3 5 8]

		* [Using cache]
		> Fibonacci-env check finished
		! Using long printer:
		length=7 capacity=8 Result: [0 1 1 2 3 5 8]

		> Required sequence part counted
		> Fibonacci-env check finished
		! Using long printer:
		length=20 capacity=32 Result: [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181]

		* [Using cache]
		> Fibonacci-env check finished
		! Using long printer:
		length=20 capacity=32 Result: [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181]
		<< Experimental part ended

		<< Main program exited

		Process finished with exit code 0
		*/
	}
}
