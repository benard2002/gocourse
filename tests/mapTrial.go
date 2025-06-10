package main

import "fmt"

func main(){

	students :=  make(map[string]int{
		"kofi" : 49,
		"Yaw": 9
	})

	fmt.Print(students)

}