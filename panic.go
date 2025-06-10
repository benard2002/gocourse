package main

import "fmt"

func main() {

	// panic(interface{})

	// Example of a valid input
	// process(10)

	// Example of invalid input
	process(-3)

}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before panic")
		panic("input nust be a non-negative number")
		// fmt.Println("After Panic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}

// panic in go is a built in funciton that stops normal execution of a function immediately.
// when a function encounters a panic it stops executing its current activities, unwinds the stack, and then executes any deferred funcitons
// a panic is typically used ot signal an unexpected error contdition where the program cannot proceed safely.
//  interface meanns you can input any value of any type as an argument for this function
//  interface is like type any in other languages

// deffer executes even when function is panicking
// anything after panic would not be reached by the runtime

// panic after deferred functions
// in conclusion , understanding how panic works and its implications is crucial for wiritng robust go programs
// misuse of panic can lead to mis beahivours
