package main

import "fmt"

var z int = 50

// Create your own type as hotdog
type hotdog int

var a hotdog

func main() {
	a = 20

	// Print z and it's type
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	// Print z and it's type
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// You can't assign value with diff type
	// z = a
	// This statement  will throw exception because int and hotdog is diff type
	// To resolve this error we need to do type conversion
	z = int(a)
	// after converting a to int it will allow to assign because type of z is int
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
