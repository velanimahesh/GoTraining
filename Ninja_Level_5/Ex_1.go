// Create struct of 3 elements name, age and fav sport
// Where fav sport consist of an list of sport name

package main

import "fmt"

type person struct {
	name     string
	age      int
	favSport []string
}

func main() {

	person1 := person{
		name: "Mahesh",
		age:  25,
		favSport: []string{
			"Cricket",
			"football",
			"carrom",
		},
	}

	person2 := person{
		name: "Dhyey",
		age:  4,
		favSport: []string{
			"Chess",
			"VolleyBall",
			"Tennis",
		},
	}

	fmt.Println(person1)
	fmt.Printf("%T\n", person1)
	fmt.Println(person1.favSport)

	fmt.Println(person2)
	fmt.Printf("%T\n", person2)
	fmt.Println(person2.favSport)

	// Iterate over sports
	for index, value := range person1.favSport {
		fmt.Println("Index = ", index, " value = ", value)
	}
	for index, value := range person2.favSport {
		fmt.Println("Index = ", index, " value = ", value)
	}

}
