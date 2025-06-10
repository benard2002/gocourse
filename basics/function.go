package main

import "fmt"

func main() {

	// func <name> (parameter list) returnType{
	// code block

	// return value
	//}

	// public function should start with an uppercase and private lowercase
	// eg. fmt.Println()

	sum := add(1, 2)
	fmt.Println(sum)

	fmt.Println(add(2, 3))

	// anonymous function
	func() {
		fmt.Println("Hello Anonymous Function")
	}() // the function is called i mmediately

	greet := func() {
		fmt.Println("Hello anonymous Function")
	}

	greet()

	// function as types
	operation := add
	result := operation(4, 5)

	fmt.Println(result)
}
func add(a, b int) int {

	return a + b
}
