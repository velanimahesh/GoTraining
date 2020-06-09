// From the previous example craete dict storing two diff values

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
	fmt.Println(person2)

	// Create dict
	dict1 := map[string]person{
		person1.name: person1,
		person2.name: person2,
	}

	for key, val := range dict1 {
		fmt.Println("key = ", key, " val = ", val.favSport)
	}
}
