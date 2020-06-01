// Create own type and assign value to var with own type

package main

import "fmt"

type own_type int
var x own_type

func main()  {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
