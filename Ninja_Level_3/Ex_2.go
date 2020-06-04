// Print every upper case 3 times
package main

import "fmt"

func main()  {

	decimalVal := 65

	//  This type of for loop is a replacement of while loop in go
	// In go there is no do_while loop
	for decimalVal <= 90 {
		for i := 0; i < 3; i++ {
			// This used to convert decimal into ascii
			fmt.Printf("%#U\n", decimalVal)
		}
		fmt.Println()
		decimalVal++
	}

}
