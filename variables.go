package main

import "fmt"

// Define and initialize global variable
var y = 10

// Define var with default value
//int  = 0
//float = 0.0
//string = ""
//bool = false
//nil : same as a null
var z int

func main() {

	// Print global variable y
	fmt.Println(y)

	// Print global var z
	fmt.Println(z)

	// define and print local variables
	// this called a short declaration operator
	x := 20
	fmt.Println(x)

	// update the value of z and y
	y = 99
	z = 100
	foo()
}

func foo() {
	// here y = 99
	fmt.Println("Foo : y = ", y)
	// here z = 100
	fmt.Println("Foo : z = ", z)
}
