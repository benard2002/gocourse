package basics

import "fmt"

func main() {

	//  in go there is no while the for is used

	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Iteration:", i)
	// 	i++
	// }

	// infinite
	// for {
	// 	fmt.Println("Hello")
	// }

	// for as while with break
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	// continue

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number", num)

		num++
	}
}
