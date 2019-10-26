package main

import (
	"fmt"
)

var truncate float64

func main() {
	fmt.Printf("Truncate Assessment:\n Enter Float Number to truncate it \n Enter a letter to stop programm:\n")
	fmt.Printf("Please enter your Float: ")
	var check = true
	for check == true {
		_, err := fmt.Scan(&truncate)
		if err != nil {
			check = false;
			continue;
		}
		fmt.Printf("Your number is: %v\n", int(truncate))
		fmt.Printf("Enter your next Float: ")
	}
	fmt.Printf("Program stopped\n")
}