package main

import "fmt"

// Separate declaration of constants
const a int  = 10
const b float64  = 20.45
const c string  = "James Bond"


// Multiple constant in same defination
const (
	x int = 20
	y string = "Hello world!!!"
)

func main()  {
	fmt.Println(a, b, c)
	fmt.Println(x, y)

	// update the value of constant
	// a = 40
	// fmt.Println(a)
	// This statement will throw error k can't assign  value to the constant
}


