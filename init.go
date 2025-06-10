package main

import "fmt"

func init() {
	fmt.Println("Initializing package...")
}
func init() {
	fmt.Println("Initializing package2...")
}
func init() {
	fmt.Println("Initializing package3...")
}
func main() {

	fmt.Println("Inside the main funciton")
}

// init function is used to perform initialization task for the package before its used
// it has no parameters and no return values
// it always get executed befor the main function

// Practical Use Cases
// Setup Tasks
// Configuration
// Registering Components
// Database Initialization
// Best Practices
// Avoid Side Effects
// Initialization Order
// Documentation
