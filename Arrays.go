package main

import "fmt"

// Declare Global array of 5 elements
// Type must require for array and all elements inside the
// array must be a same type
var arr_1 [5]int
var arr_4 []int

func main() {

	// Print array elements
	// By default all the elements inside array is 0
	fmt.Println(arr_1) // [0 0 0 0 0]

	// Assign value to array by indexing
	arr_1[0] = 42 // Indexing start with 0 same as python
	arr_1[3] = 43
	fmt.Println(arr_1) // [42 0 0 43 0]
	// arr_1[5] = 100    // Exception:  invalid array index 5 (out of bounds for 5-element array

	// Get length of an array
	fmt.Println(len(arr_1)) // 5

	// Declare local array
	// Init array during declaration also called a slices
	arr_2 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr_2) // [0 1 2 3 4 5 6 7 8 9]

	// By default array put 0 values to non-assign elements
	arr_3 := [10]int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr_3) // [0 1 2 3 4 5 0 0 0 0]

	// Print size less array
	fmt.Println(arr_4)      // []
	fmt.Println(len(arr_4)) // 0

	// ITERATE over ARRAY ELEMENTS
	arr_5 := [10]int{9, 5, 2, 3, 4, 2, 6, 8, 2, 1}
	fmt.Println(arr_5) // [0 1 2 3 4 5 6 7 8 9]
	for index, value := range arr_5 {
		fmt.Println("Index = ", index, "value = ", value)
	}

	// Slices in Go
	var arr_6 [4]int
	arr_6[0] = 5
	arr_6[1] = 6
	arr_6[2] = 7
	arr_6[3] = 8
	// Print all element using slice
	fmt.Println(arr_6[:]) // [5 6 7 8]
	// print element 1 & 2 only
	fmt.Println(arr_6[:2]) // [5 6]
	// Print last element only
	fmt.Println(arr_6[2:]) // [7 8]

	// Append item into an Array
	// Array must be empty for the append action
	arr_7 := []int{3, 4, 5, 6, 7}
	fmt.Println(arr_7)
	arr_7 = append(arr_7, 8, 9, 10)
	fmt.Println(arr_7) // [3 4 5 6 7 8 9 10]

	// Append Array into an Array
	arr_8 := []int{3, 4, 5, 6, 7}
	arr_9 := []int{8, 9, 10}
	arr_8 = append(arr_8, arr_9...) // 3 dots here indicate append all element for 2nd array
	fmt.Println(arr_8)              // [3 4 5 6 7 8 9 10]

	// Delete value 4 from Array
	arr_10 := []int{3, 4, 5, 6, 7}
	arr_10 = append(arr_10[:1], arr_10[2:]...)
	fmt.Println(arr_10) // [3 5 6 7]

	// MAKE in slice/array
	//make (type of an array, length of array, capacity of array
	x := make([]int, 10, 15)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(x)
	fmt.Println(len(x)) // 20
	fmt.Println(cap(x)) // 30  Not cleared

	// MultiDimenssional Array
	p := []string{"a", "b", "c", "d"}
	fmt.Println(p)
	q := []string{"w", "x", "y", "z"}
	fmt.Println(q)
	r := [][]string{p, q}
	fmt.Println(r) // [[a b c d] [w x y z]]
}
