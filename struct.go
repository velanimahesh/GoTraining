package main

import "fmt"

// Define a Structure
type person struct {
	first_name string
	last_name  string
}

type secretAgent struct {
	person
	licenceTooKill bool
}

func main() {

	p1 := person{
		first_name: "Mahesh",
		last_name:  "Velani",
	}

	p2 := person{
		first_name: "Dhyey",
		last_name:  "Patel",
	}

	// Access struct element
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first_name, " ", p1.last_name)
	fmt.Println(p2.first_name, " ", p2.last_name)

	// Embedding the struct element
	s1 := secretAgent{
		person: person{
			first_name: "Mahesh",
			last_name:  "Velani",
		},
		licenceTooKill: true,
	}
	fmt.Println(s1)
	fmt.Println(s1.first_name)
	fmt.Println(s1.licenceTooKill)

	// Anonymous struct (A struct without Name)
	p3 := struct {
		first_name string
		last_name  string
		age        int
	}{
		first_name: "AnonymousFirst",
		last_name:  "AnonymousLast",
		age:        32,
	}
	fmt.Println(p3)
}
