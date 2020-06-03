// iota is nothing bt ENUM where once we define it it will auto increase the vales

package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d = iota
)

// Also define as
const (
	x = iota
	y
	z = iota + 10  // We can also perform expression with iota E.g: here o/p will be 2 + 10 = 12
)

func main()  {
	fmt.Println(a, b, c, d)  // O/p = 0 1 2 3
	fmt.Println(x, y, z)  // O/p = 0 1 12
}
