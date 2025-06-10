package basics

import "fmt"

func main() {

	// var arryName [size]elementType
	var numbers [5]int
	fmt.Println(numbers)

	// adding values into the array
	// fmt.Print(len(numbers)) the length of the array

	numbers[0] = 9
	numbers[4] = 20

	fmt.Print(numbers)

	// initializing array with value

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array: ", fruits)

	fmt.Println("Third element: ", fruits[2])

	nums := [5]int{1, 2, 3, 4, 5}
	// iterating over the elements of an array
	// for i := 0; i < len(nums); i++ {

	// 	fmt.Println("Element at index, ", i, ": ", nums[i])
	// }

	// another way is to use the range function

	// for index, value := range nums {

	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)

	// }

	// for i, v := range nums {

	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)

	// }

	// discarding i
	// underscore is blank identifier, used to store unused values
	for _, v := range nums {
		fmt.Printf("Value: %d\n", v)
	}

	// a, b := someFunction()
	// fmt.Println(a)
	// fmt.Println(b)

	// if we dont want to use b we can replace with an underscore (_)

	a, _ := someFunction()
	fmt.Println(a)

	//
	b := 2
	_ = b // for variable we're not usign temporarily

	// lenght of an array
	// fmt.Println("The length of nums array is ", len(nums))

	// // comparing arrays
	// array1 := [3]int{1, 2, 3}
	// array2 := [3]int{1, 2, 3}

	// fmt.Println("Array1 is equall to array2", array1 == array2)

	// // multidimentional Array
	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(matrix)'

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println(*copiedArray, originalArray)
}

func someFunction() (int, int) {
	return 1, 2
}
