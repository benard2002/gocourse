package basics

import "fmt"

func main() {
	// variable declaration
	var a, b int = 10, 3

	var result int

	result = a + b
	fmt.Println("Addition", result)

	result = a - b
	fmt.Println("Subtraction", result)

	result = a * b
	fmt.Println("Multiplication", result)

	result = a % b
	fmt.Println("Remainder", result)

	result = a / b
	fmt.Println("Division", result) // dividing two integers gives an interger

	const p float64 = 22 / 7.0 // execution starts from the right to the left
	const k = 22.0 / 7.0       // works fine
	fmt.Println(p, k)

	// overflow and underflow

	// overflow with signed integers

}
