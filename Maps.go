package main

import "fmt"

func main() {

	// Create a dictionary store name and Age values
	map_1 := map[string]int{
		"Mahesh": 26,
		"Dhyey":  4,
	}
	fmt.Println(map_1)           // map[mahesh:26 Dhyey:4]
	fmt.Println(map_1["Mahesh"]) // 26
	fmt.Println(map_1["Dhyey"])  // 4
	fmt.Println(map_1["Ajay"])   // 0  Print val = 0 is no key present

	// Check key is present or not
	if v, ok := map_1["Ajay"]; ok {
		fmt.Println("Ajay is present in list", v)
	}
	if v, ok := map_1["Mahesh"]; ok {
		fmt.Println("Mahesh is present in list", v)
	}

	// Add new value to map and iterate over loop
	map_1["Ajay"] = 21
	fmt.Println(map_1)
	// Iterate over mapping elements
	for key, val := range map_1 {
		fmt.Println("Name = ", key, "& Age = ", val)
	}

	// Delete an entry from map
	delete(map_1, "Ajay") // Delete Ajay Entry from the mapping list
	fmt.Println(map_1)    // map[Dhyey:4 Mahesh:26]
	// Check value deleted
	if v, ok := map_1["Ajay"]; ok {
		fmt.Println("Ajay is present in map list ", v)
	} else {
		fmt.Println("Ajay key is deleted from the table")
	}
}
