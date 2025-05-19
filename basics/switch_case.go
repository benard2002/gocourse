package main

import "fmt"

func main() {

	// switch expression {
	// case value1:
	// something to do for first condition
	// break
	// but note that go automatically breaks ones a condition is met
	// use fallthrough to indicate a continuation to the next case as soon as a condition is met
	//  eg cas value1 (fallthrough):

	// case value2:
	// code to be executed if expression equals value2

	// case value3:
	// code to be executed if expression equal value3

	// default:
	//  code to be executed if expression does not match any value
	// break // if you want

	// }

	// fruit := "apple"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple.")

	// case "banana":
	// 	fmt.Print("It's a banana.")

	// default:
	// 	fmt.Print("Unknown fruit.")

	// }

	// // multiple cases
	// day := "MOnday"

	// switch day {

	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")

	// case "Saturday", "Sunday":
	// 	fmt.Println("It's a  weekend")

	// default:
	// 	fmt.Println("Invalid day")

	// }

	// switch with fallthrough

	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")

	// default:
	// 	fmt.Println("Not Two")

	// }

	checkType(10)
	checkType(3.14)
	checkType("Hell")
	checkType(true)
}

func checkType(x any) {
	// x interface => that is for any datatypes

	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknwown Type")
	}
}
