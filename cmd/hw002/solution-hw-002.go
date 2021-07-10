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
	"github.com/rodkevich/go-course/homework/hw002/fibo"
	"os"
)

//appExit exit app with error print on wrong use
func appExit(code int, s string) {
	fmt.Println(s)
	os.Exit(code)
}

//hwPrinterSolution  prints a sequence of fibonacci numbers up to required length
func hwPrinterSolution(n int) {
	fmt.Println("Solution part started...")
	var sequence, err = fibo.SizedSequence64bit(int64(n))
	if err != nil {
		appExit(1, err.Error())
	}
	fmt.Println("Result:")
	fmt.Println(sequence)
}

func main() {
	defer fmt.Println("\n...Main program exited")
	experimentalPart := flag.Bool("e", false, "Run with experiments")
	flag.Parse()
	var n int
	fmt.Print("Enter an integer between 1 - 93: \n")
	_, err := fmt.Scanln(&n) // get input value into n var
	if err != nil {
		appExit(1, err.Error())
	}
	//Solution part
	hwPrinterSolution(n)

	//Experimental part
	if *experimentalPart  {
		hwOptionalExperiments(n)
	}
}

// hwPrinterSolution(n) output for n = 25:
/*
Enter an integer between 1 - 93:
25
Solution part started...

Counting fresh sequence finished
Result:
[0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368]

...Main program exited
*/

//☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣
////////////////////  Experimental   ///////////////////////
////////////////////      part       ///////////////////////
//☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣☣

func hwOptionalExperiments(n int) {
	fmt.Println("\nExperimental part started...")
	cache := fibo.NewCache()
	printSequenceUsingCache(int64(n), &cache)
	printSequenceUsingCache(7, &cache)
	printSequenceUsingCache(4, &cache)
	printSequenceUsingCache(20, &cache)
	printSequenceUsingCache(17, &cache)

}

func printSequenceUsingCache(n int64, cache *fibo.Cache) []int64 {
	var sequence, err = fibo.SizedSequence64bitUsingCache(n, cache)
	if err != nil {
		appExit(1, err.Error())
	}
	fmt.Println("Result:")
	fmt.Println(sequence)
	return sequence
}

// hwPrinterSolution(n) +  hwOptionalExperiments(n) output for n = 5:
/*
Enter an integer between 1 - 93:
5
Solution part started...

Counting fresh sequence finished
Result:
[0 1 1 2 3]

Experimental part started...

Counting fresh sequence finished
Result:
[0 1 1 2 3]

Counting fresh sequence finished
Result:
[0 1 1 2 3 5 8]

Using cache
Result:
[0 1 1 2]

Counting fresh sequence finished
Result:
[0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181]

Using cache
Result:
[0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987]

...Main program exited

Process finished with exit code 0
*/
