package main

import (
	"fmt"
)

func main() {
	var input float64

	fmt.Printf("Please enter a floating point (with decimals) number: ")

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Printf("There was an error: %s", err)
	} else {
		// fmt.Printf("%d \n", int(entered))
		fmt.Printf("Here is your floating point number, truncated to an int: %v.\n", int(input))
	}

}
