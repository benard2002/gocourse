package basics

import "fmt"

func main() {

	// temperature := 25

	// if temperature >= 30 {
	// 	fmt.Println("It's hot outside!")
	// } else {
	// 	fmt.Println("It's cool outside!")
	// }

	// score := 10
	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade C")
	// } else {
	// 	fmt.Println("Grade D")
	// }

	// num := 30

	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println("Number is divisible by both 2 and 3.")
	// 	} else {
	// 		fmt.Println("Number is divisible by 2 but not 3.")
	// 	}
	// } else {
	// 	fmt.Println("Number is not divisible by 2")
	// }

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("Either 10 or 5 is even.")
	}
	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 and 6 are even.")
	}

}
