package main

import "fmt"

func main() {
	process()
	fmt.Println("Returned from process")

}

func process() {
	defer func() {
		// if r := recover(); r != nil { // the code is just like declaring r:= recover and writing below it if r!= nil
		// 	fmt.Println("Recovered:", r)
		// }
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}

// recover is a built in function used to regain control of a panicking go routine
// its only useful inside defer funcitons and is used to manage behavior of a panicking go routine to avaoid abrupt termination
// recover is directly related to panic , its used to regain control of a panicking go routine
// recover can be used to handle or log errors gracefully and allow the program to continue running

// the recover function is called in a defer function
// recover() is used inside a defer function.

// It catches the panic, so your program can recover and keep running
// recover(): built-in function allow a program to gain control after a panic, recover returns nil if there is no panic

// Practical Use Cases
// Graceful recovery
// cleanup
// logging and reporting
// Best Practices
// Use with Recovery
// Avoid silent Recovery
// Avoid overuse
