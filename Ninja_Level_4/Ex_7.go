// Multi dimension array

package main

import "fmt"

func main() {
	x := []string{"1", "2", "3", "4"}
	y := []string{"A", "B", "C", "D"}

	z := [][]string{x, y}
	fmt.Println(z)
}
