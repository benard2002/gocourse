package main

import "fmt"

func main() {
	// var sliceName[] ElementType

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}
	// numbers2 := []int{9, 8, 7}

	// using the make function to set fixed size
	// slice := make([]int, 5)

	// fmt.Println(slice)

	// slicing an array
	a := [5]int{1, 2, 3, 4, 5} // an array,ie: a fixed size is set
	slice1 := a[1:4]           // from index1 of array a to 4
	// slice2 := a[1:]            // from index 1 to last

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// appending into a slice
	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)

	// making a twoD
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
