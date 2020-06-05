// Create a 5 element int array, assign value by index
// Iterate over element via range and print type of array

package main

import "fmt"

var arr [5]int

func main() {

	arr[0] = 9
	arr[1] = 8
	arr[2] = 7
	arr[3] = 6
	arr[4] = 5
	fmt.Println(arr)

	// Also do like
	//arr := []int{9,8,7,6,5}
	//fmt.Println(arr)

	for index, value := range arr {
		println("ndex = ", index, " & value = ", value)
	}

	fmt.Printf("Type of an ARRAY = %T\n", arr)
}
