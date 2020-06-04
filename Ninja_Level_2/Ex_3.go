// Create type and un-type constants

package main

import "fmt"

const (
	a int = 2         // Type constants
	b = "James Bond"  // un-type constants
)

func main()  {

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", b)
}