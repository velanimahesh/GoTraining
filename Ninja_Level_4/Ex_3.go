// Create Slicing of diff len from previous example

package main

import "fmt"

func main() {

	// Define slice
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)

	// Print 2 to 5
	arr_1 := arr[1:5]
	fmt.Println(arr_1)

	// Print 1 to 5
	arr_2 := arr[:5]
	fmt.Println(arr_2)

	// Print 5 to 10
	arr_3 := arr[5:]
	fmt.Println(arr_3)

}
