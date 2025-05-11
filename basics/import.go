package basics
// import "fmt" for single import package
// for multiple packages
import (
	"fmt"
	"net/http"
)
// named import
import foo "net/http" 
// that means we use foo inplace of the import package

func main()  {
	fmt.Println("Hell world")
}