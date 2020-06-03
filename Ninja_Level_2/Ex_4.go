// Bit Shifting

package main

import "fmt"

var z int = 2

func main()  {
	fmt.Printf("Decimal = %d\n", z)
	fmt.Printf("Bin = %b\n", z)
	fmt.Printf("Hex = %#x\n", z)
	fmt.Printf("Int = %v\n", z)

	// Shift left
	x := z << 1
	fmt.Printf("Decimal = %d\n", x)
	fmt.Printf("Bin = %b\n", x)
	fmt.Printf("Hex = %#x\n", x)
	fmt.Printf("Int = %v\n", x)

	// Shift Right
	y := z >> 1
	fmt.Printf("Decimal = %d\n", y)
	fmt.Printf("Bin = %b\n", y)
	fmt.Printf("Hex = %#x\n", y)
	fmt.Printf("Int = %v\n", y)

}
