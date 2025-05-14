package main

import "fmt"

var middleName = "Bobo"

func main() {

	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	middleName = "George"
	// count := 10
	// lastName := "Seth"

	fmt.Println(middleName)
}

// scope

func Printname() {
	firstName := "Ben"

	fmt.Println(firstName)
	middleName = "Bro code"
}
