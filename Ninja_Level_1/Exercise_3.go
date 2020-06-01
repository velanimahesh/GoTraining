// At package level scope assign the values
// Use sprintf to print the values and assign to variable s
// also print the s variable

package main
import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main()  {
	s := fmt.Sprintf("x = %v\ty = %v\tz = %v\n", x, y, z)
	fmt.Println(s)
}