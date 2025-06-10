package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Deferred statement")

	fmt.Println("Starting the main function")
	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of main function")
}

// os.exit is a function tha exit the program immediately with some code
// any defer funciton is not executed
// the function take an integer argument rep the status code returned to the operating system
// status code indicates successful completion , any non zoero indicates and error or abnormal termination
// os.exit bypasses the normal defer, panic , and recover mechanisms

// Practical Use Cases
// Error Handling
// Termination Conditions
// Exit Codes

// Best Practices
// Avoid Deferred Actions
// Status Codes
// Avoid Abusive Use
