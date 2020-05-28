package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("I am executing 1st func...")
	firstFunc()
	fmt.Println("I am executing loop func...")
	funcLoop()
	fmt.Println("Exiting from the program")
}

func firstFunc() {
	fmt.Println("I am inside 1st sample func")
}

func funcLoop() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
