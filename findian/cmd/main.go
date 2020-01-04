package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Find in String Assessment:\nEnter a String to start search for: i a n \nEnter 'exit' to stop programm:\n")
	var check = true
	var find = bufio.NewScanner(os.Stdin)
	for check == true {
		fmt.Printf("Please enter your String: ")
		find.Scan()
		s := strings.ToLower(find.Text())
		length := len(s)
		// check if user wants to exit program
		if strings.Compare(s, "exit") == 0 {
			break
		}
		// check if i is first charakter
		if strings.Index(s, "i") != 0 {
			fmt.Printf("Not Found!\n")
			continue
		}
		// check if a is somewhere
		if strings.Index(s, "a") == -1 {
			fmt.Printf("Not Found!\n")
			continue
		}
		// check if n is at the end
		if strings.Index(s, "n") != length-1 {
			fmt.Printf("Not Found!\n")
			continue
		}
		fmt.Printf("Found!\n")
	}
	fmt.Printf("Program stopped\n")
}