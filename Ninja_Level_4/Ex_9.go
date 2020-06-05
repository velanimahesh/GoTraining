// Add one more record in previous example
// Also delete newly added entry from the dict

package main

import "fmt"

func main() {
	// List as a dictionary value
	map_2 := map[string][]string{
		"Mahesh": []string{"Coding", "sleeping", "reading"},
		"Dhyey":  []string{"playing", "crying", "video Gaming"},
		"Ajay":   []string{"PUBG", "Teen Patti", "TikTok"},
	}

	for key, val := range map_2 {
		fmt.Println("key = ", key)
		for index, value := range val {
			fmt.Println("index = ", index, " & val = ", value)
		}
	}

	// Add one more entry into the dict
	map_2["John Cena"] = []string{"Wrestling", "Karate", "Eating"}
	fmt.Println(map_2)
	fmt.Println(len(map_2))

	// Delete entry of the John from the dict
	delete(map_2, "John Cena")
	fmt.Println(map_2)
	fmt.Println(len(map_2))

}
