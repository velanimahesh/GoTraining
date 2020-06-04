// Switch with expression
package main

import "fmt"

var favSport string = "Cricket"

func main()  {
	switch favSport {
	case "Cricket":
		fmt.Println("My Fav Sport is Cricket")
	case "Football":
		fmt.Println("My Fav Sport is Football")
	default:
		fmt.Println("I don't liking a Cricket or Football")
	}
}
