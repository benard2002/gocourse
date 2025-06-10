package basics

import "fmt"

func main() {
	// Maps

	// var mapVariable map[keyType]ValueType

	// mapVariable = map[keyType]ValueType

	// Using a Map Literal
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)

	myMap["key1"] = 9
	myMap["Code"] = 10

	fmt.Println(myMap)

	map2 := make([string]string{
		"name": "Bro"
	})

}
