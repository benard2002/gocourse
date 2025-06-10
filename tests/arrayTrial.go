package main

import "fmt"

func main() {

	integers := [5]int{3, 5, 7, 9, 11}

	sum := 0

	for _, num := range integers {
		sum += num
	}

	fmt.Printf("The sum of the integers is: %d", sum)

}
