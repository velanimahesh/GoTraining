// Map example store person as key and fav as val
// Nested dictionary

package main

import "fmt"

func main() {
	map_1 := map[string]string{
		"Mahesh": "Coding",
		"Dhyey":  "playing",
		"Ajay":   "PUBG",
	}
	fmt.Println(map_1)

	for key, val := range map_1 {
		fmt.Println(key, " Likes = ", val)
	}

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
}
