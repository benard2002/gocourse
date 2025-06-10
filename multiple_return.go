package main

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2...) (returnType1, returnType2,...){
	//code block
	//return returnvalue1, returnValue2....

	//}

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

// multiple return values helps in error handling
// rem: %v is general for any value
// use nil for what you dont have to return, nb: you cannot use nil as sting instead use empty ""
// errors
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater thatn b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
