// Anonymous Struct

package main

import "fmt"

func main() {

	// Defined anonymous struct
	person := struct {
		name string
		age  int
	}{
		name: "Mahesh Velani",
		age:  25,
	}

	fmt.Println(person)
}
