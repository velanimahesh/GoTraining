// Create own type and assign value to var with own type

package main

import "fmt"

type own_type int
var x own_type
var y int            // Here int is called a underlying type because own_type defined from  it.

func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)

	// Assign own_type value to y
	y := int(x)                     // Conversion require because type of x is own_type
	fmt.Println("y = ", y)
	fmt.Printf("%T\n", y)
}
