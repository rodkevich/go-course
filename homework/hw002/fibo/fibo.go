package fibo

import (
	"errors"
	"fmt"
	"os"
)

var (
	//error to try customizing
	errorLimitOf64bitSeq = errors.New("\nAbort app: wrong sequence size provided (max 64bit sequence size is 93)")
	errorBadEnvVars      = errors.New("\nAbort app: required env vars were not correct")
	errorNotImplemented  = errors.New("\nAbort app: implementation is not there yet. See ya!")
	knownPrinters        = []struct {
		name        string
		implemented bool
	}{{"long", true}, {"short", false}}
)

const limitOf64bitSeq int64 = 93 //limit sequence length with a const for validation

// init function prints message to show when Go runs this kind for packages
func init() {
	fmt.Println("setup: Fibonacci-app initialized")
}

//checkEnv experiment to work with basic types
// function sets env vars to refer in printer function
func checkEnv() error {
	defer fmt.Println("> Fibonacci-env check finished")
	var canPrinterBeUsed bool //default `false`
	var printerName = os.Getenv("Printer")
	//define if required printer can be used
	for _, printer := range knownPrinters {
		if printer.name == printerName {
			//set true/false
			canPrinterBeUsed = printer.implemented
		}
	}
	//
	if canPrinterBeUsed == false {
		err := os.Setenv("Printer", "notAllowed")
		if err != nil {
			return errorBadEnvVars
		}
	}
	return nil
}

//Solution prints a sequence of fibonacci numbers up to required length
func Solution(n int) error {
	defer fmt.Println("<< Solution part ended")
	var sequence, err = sizedSequence64bit(int64(n))
	if err != nil {
		return err
	}
	err = printSlice(sequence)
	if err != nil {
		return err
	}
	return nil
}

// sizedSequence64bit returns a new slice of fibonacci numbers
//with required length if possible. Breaks on 64 bit overflow
func sizedSequence64bit(size int64) ([]int64, error) {

	var (
		counted int64
		a       int64   = 0
		b       int64   = 1
		rtn     []int64 = []int64{a, b} // init with first two
	)
	//Check input argument values to be expected. If not -> return
	if size > limitOf64bitSeq {
		// TODO: почитать про nil
		// TODO: почитать type assertion
		return nil, errorLimitOf64bitSeq
	}
	for i := int64(2); i <= size; i++ {
		counted = a + b
		//experimental checkEnv if max value reached because input is already validated
		if counted < 0 {
			fmt.Printf("! Warn: 64 bit overflow on %d iteration", i+1)
			size = i
			break
		} else {
			//prepare vars for new iteration
			a, b = b, counted
			rtn = append(rtn, counted)
		}
	}
	fmt.Println("\n> Required sequence part counted")
	return rtn[0:size], nil
}

//////////////////  Experimental   /////////////////////
//////////////////      part       /////////////////////

// Cache custom type to play with
type Cache []int64

// NewCache function returns prepared cache with first two numbers
func NewCache() Cache {
	return Cache{0, 1}
}

func OptionalSolutionWithCaching(n int64, cache *Cache) error {
	var sequence, err = sizedSequence64bitWithCache(n, cache)
	if err != nil {
		return err
	}
	err = printSlice(sequence)
	if err != nil {
		return errorBadEnvVars
	}
	return nil
}

// sizedSequence64bitWithCache uses cache or returns a new slice of fibonacci numbers
//with required length if possible. Breaks on 64 bit overflow
func sizedSequence64bitWithCache(size int64, cache *Cache) ([]int64, error) {
	if size > int64(len(*cache)) { // convert type to be comparable
		out, err := sizedSequence64bit(size)
		if err != nil {
			return nil, err
		}
		*cache = out
		return out, err
	} else {
		fmt.Println("\n* [Using cache]")
		////////////////////////////
		// some experiments to play
		rtn := make([]int64, 0)
		for _, v := range *cache {
			rtn = append(rtn, v)
		}
		return rtn, nil
	}
}

func printSlice(s []int64) error {
	err := checkEnv() //check env vars to be OK
	if err != nil {
		return errorBadEnvVars
	}
	//get required printer from env
	fmt.Printf("! Using %v printer:\n", os.Getenv("Printer"))
	switch printer := os.Getenv("Printer"); printer {
	case "long":
		fmt.Printf("length=%d capacity=%d Result: %v\n", len(s), cap(s), s)
	case "short":
		fmt.Println(s)
	case "notAllowed":
		fmt.Println("! Selected printer is not implemented")
		return errorNotImplemented
	}
	return nil
}
