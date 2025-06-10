package main

import "fmt"

func main() {

	process(10)
}

// in go defer is a mechanism that allows you to postpone the execution of a function untill the surrounding function returns ,
//  its useful feature to ensure that certain cleanup actions or finalizing tasks are performed,
// regardless of how the function exists whether it returns normally or panics.
// multiple defer statements are executed in a last in first out order when the function returns
// arguments to differed functions are evaluated immediately when the differ statement is encountered.
func process(i int) {
	fmt.Println("Deferred i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i: ", i)
}

// Practical Use Cases
// Resource Cleanup
// Unlocking Mutexes
// Logging and Traicing

// Best Practices
//Keep Deferred Actions Short
//Understand Evaluation Timing
//Avoid Complex Control Flow
