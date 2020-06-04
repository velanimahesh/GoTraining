// Switch without expression
package main

import "fmt"

func main()  {
	x := 5
	switch {
	case x < 5:
		fmt.Println("number is less then 5")
	case x == 5:
		fmt.Println("number is 5")
	default:
		fmt.Println("Number is greater then 5")
	}
}
