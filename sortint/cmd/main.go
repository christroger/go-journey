package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	fmt.Printf("Sort Integer:\nEnter a Integer to sort it in Slice \nEnter 'X' to stop programm:\n")
	var list = make([]int, 3)
	var check = true
	var find = bufio.NewScanner(os.Stdin)
	var notFull = true
	var controll []int
	for check == true {
		fmt.Printf("Please enter your Integer: ")
		find.Scan()
		t := find.Text()
		// check if user wants to exit program
		if strings.Compare(t, "X") == 0 {
			break
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("That was not an Integer nor X")
			continue
		}
		if notFull == false {
			list = append(list, i)
		} else {
			list[0] = i
		}
		controll = append(controll, i)
		if len(controll) == 3 {
			notFull = false
		}

		sort.Ints(list)
		fmt.Printf("Sorted List: %v\n", list)

	}
	fmt.Printf("Program stopped\n")
}