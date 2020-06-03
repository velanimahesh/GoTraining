// Assign value using different operator

package main

import "fmt"

func main()  {

	a := (42 == 42)
	b := (42 <= 41)
	c := (42 >= 41)
	d := (42 < 40)
	e := (42 > 40)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)

}