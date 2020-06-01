// Use Short declaration Operators to
// Assign value to identifier x, y and z
//values = 42, "James Bond", true

// Print value stored in this vars using single and multiple print statement
package main

import "fmt"

func main()  {
	x := 42
	y := "James Bond"
	z := true

	// Using single print statement
	fmt.Println("x = ", x, "\ty = ", y, "\tz = ", z)

	// Using multiple print statement
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	fmt.Printf("z = %v\n", z)
}