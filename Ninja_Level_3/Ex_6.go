// if_else and else_if

package main

import "fmt"

func main()  {
	x := 4
	y := 7
	z := 7

	// Satisfy if condition
	if x < 5{
		fmt.Println("X Number is less than 5")
	} else {
		fmt.Println("X Number is greater than 5")
	}

	// Fulfil Else condition
	if y < 5{
		fmt.Println("Y Number is less than 5")
	} else {
		fmt.Println("Y Number is greater than 5")
	}

	// Satisfy else_if condition
	if z < 5{
		fmt.Println("Z Number is less than 5")
	} else if z < 10 {
		fmt.Println("Z Number is between 5 to 10")
	} else {
		fmt.Println("Z is greater than 10")
	}
}
