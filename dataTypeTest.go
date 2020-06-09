package main

import "fmt"

// Go is a static language so specific variable
// consist only specific type.
// e.g. variable declared/init with int consist only int value
// and you can't assign a str value to this particular variable

// Int declaration
var y = 42

// Str declaration
var z = "Hello, World!!!"

// specific declaration with type
var x int = 44

// multiline string
// you can also define double quoted statement inside this str
var a string = `I am multiline declaration,
write multiple line here,
That's it`

func main() {
	// print y
	fmt.Println(y)
	// print type of y
	fmt.Printf("%T\n", y)

	// print z
	fmt.Println(z)
	// print type of z
	fmt.Printf("%T\n", z)

	// print x
	fmt.Println(x)
	// print type of x
	fmt.Printf("%T\n", x)

	// assign int val to str type
	//z = 43
	// This statement will throw an exception that
	// you can't assign int value to str type

	// print a
	fmt.Println(a)
	// print type of a
	fmt.Printf("%T\n", a)
}
