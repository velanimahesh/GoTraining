package main

import "fmt"

func main() {
	x := 2

	fmt.Printf("%d\t%b\n", x, x) // o/p = 2  10

	// shift 1 bit left
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y) // o/p = 4  100

	// shift 1 bit right
	z := y >> 1
	fmt.Printf("%d\t%b\n", z, z) // o/p = 2  10
}
