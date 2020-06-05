// Delete value 11 to 15 fromm below slice
// [2 3 4 5 6 7 8 9 0 11 12 13 14 15 41 42 43 44 45]

package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 41, 42, 43, 44, 45}
	fmt.Println(x)

	// Delete values
	x = append(x[:9], x[14:]...)
	fmt.Println(x)

}
