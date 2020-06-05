// Create a 10 element Slice, assign value by index
// Iterate over element via range and print type of array

package main

import "fmt"

func main() {

	// Empty array called a slice
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)

	for index, val := range arr {
		fmt.Println("index = ", index, " & val = ", val)
	}

	fmt.Printf("Type of slice = %T\n", arr)
}
