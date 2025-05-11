package basics

import "fmt"

var middleName = "Cane"

func main(){

	// uninitialized variable
	// var age int 
	// initialized variable
	// var name string = "John" // type is optional as below

	// var name1 = "John"

	//count := 10 // go by this go simple standard, the type is infered by the go compiler
	// lastName := "Ben"

	// Default values
	// numeric Types: 0
	// boolean : false
	// string Type: ""
	// Pointer, Sliced, maps, functions, adn structs: nil

}

// scope
func printName(){
	firstName := "KO"
	fmt.Println(firstName)
}