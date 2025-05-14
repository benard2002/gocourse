package basics

import "fmt"

// below is struct
type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// PascalCase
	// Eg. Calculus

	// snake_case
	// Eg. user_id, first_name, http_request

	// UPPERCASE
	// use case is constant

	// mixedCase
	//EG. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001

	fmt.Println(employeeID)
	fmt.Print("EmployeeId: ", employeeID)

}
