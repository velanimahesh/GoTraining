// Append value into Slice

package main

import "fmt"

func main() {

	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)

	// Append value 11 to slice
	arr = append(arr, 11)
	fmt.Println(arr)

	// Append 12-15 into slice
	arr = append(arr, 12, 13, 14, 15)
	fmt.Println(arr)

	// Append Array into slice
	x := []int{41, 42, 43, 44, 45}
	arr = append(arr, x...)
	fmt.Println(arr)

}
