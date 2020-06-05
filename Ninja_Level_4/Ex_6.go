// Print Array element using for loop
package main

import "fmt"

func main() {
	x := []int{2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 12, 13, 14, 15, 41, 42, 43, 44, 45}
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println("index = ", i, " & Val = ", x[i])
	}
}
