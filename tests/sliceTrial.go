package main

import "fmt"

func main() {

	fruits := []string{"apple", "banana", "cherry"}

	fruits = append(fruits, "date", "elderberry")

	fmt.Println(fruits, len(fruits))
}
