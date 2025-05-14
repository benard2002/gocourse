package basics

const pi = 3.14
const GRAVITY = 9.81 // untyped constant

// declarig constant without type will be automatilly inferred by go

func main() {

	const days int = 7

	// declaring multiple constants
	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
		friday        = 5
		saturday      = 6
		sunday        = 7
	)

	// the above is called a const block

}
