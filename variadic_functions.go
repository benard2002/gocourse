package main

import "fmt"

func main() {

	// func functionName(param1 type1, param2 type2, param3 ... type3) returnType{
	// 	// function body

	// }

	// fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))
	// statement, total := sum("The sum of 1,2,3 is", 1,2,3)

	// fmt.Println(statement, total)
	// variadic functions, for loop, range are popularly used in programming in go lang

	// sequence, total := sum(1, 20, 30, 40, 50, 60)
	// fmt.Println("Sequence: ", sequence, "Total", total)

	numbers := []int{1, 2, 3, 4, 5, 9}
	// passing the numbers slice into the sum function
	sequence, total := sum(3, numbers...)
	fmt.Println("Sequence", sequence, "Total", total)

}

// variadic function to calculate the sum of integers
func sum(sequence int, nums ...int) (int, int) {

	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}

// variadic functions in go allows you to create funcitions that can accept a variable number of argument
//  In go , variadic functions are defined by prefixing the type of the last parameter with an ellipsis
// ellipsis is three dots ...
// ... Ellipsis => this indicates that function can accept zero or more arguments of that type
//  no limit ot the number of parameters you can enter for type3
//  nb: variadic parameters must always come last

// passing slices to variadic functions. you can do so by appending ellipses to the slice when calling ht afunciton and this sysntax unpac the slice to individual element
