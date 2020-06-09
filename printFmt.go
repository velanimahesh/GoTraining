package main

import "fmt"

var y = 42

////////////////////////////////////////////
// fmt.Print >> Will print without new line
// fmt.Println >> Will print with new line
// fmt.Printf >> Will print with format like type and all
/////////////////////////////////////////////

////////////////////////////////////////////
// fmt.Sprint >> Used to print string
// fmt.Sprintln >>
// fmt.Sprintf >>
/////////////////////////////////////////////

////////////////////////////////////////////
// fmt.Fprint >> Used to print file content
// fmt.Fprintln >>
// fmt.Fprintf >>
/////////////////////////////////////////////

func main() {
	// Print value in line
	fmt.Println(y)

	// print type of var
	fmt.Printf("%T\n", y)

	// Print binary val
	fmt.Printf("%b\n", y)

	// Print Hex val
	fmt.Printf("%x\n", y)

	// Print Hex val with 0x
	fmt.Printf("%#x\n", y)

	// Print Normal value
	fmt.Printf("%v\n", y)
}
