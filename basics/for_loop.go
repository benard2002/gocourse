package basics

import "fmt"

func main() {
	// 	// Simple iteration over a range
	// 	for i := 1; i <= 5; i++ {

	// 		fmt.Println(i)
	// 	}

	// iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6} // slice
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, value: %d\n", index, value)
	// 	// printf: is a command that prints according to a certain format, is to include placehoders and use variables to fill the placeholders
	// 	// %v = general value , %d = specific to numbers

	// }

	// break and continue
	// numbers := []int{1,2,3,4,5,6}
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue to loop  but skip the rest of lines/statements
	// 	}
	// 	fmt.Println("Odd Number:", i)

	// 	if i == 5 {
	// 		break // breakout of the loop
	// 	}
	// }

	// ASTERIK LAYOUT
	// rows := 5

	// // Outer loop
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}

	// 	// ineer loop for start
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}

	// 	fmt.Println() // MOve to the next line
	// }

	for i := range 10 {
		fmt.Println(i)
	}

	fmt.Println("We have a lift off!")
}
