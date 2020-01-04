package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Printf("Bubble Sort Assessment: \nEnter up to 10 Numbers or quit entering numbers by typing \"x\"\n")
	var scanner = bufio.NewScanner(os.Stdin)
	var intArray = make([]int, 0)
	
	for x := 0; x < 10; x++ {
		fmt.Printf("Please Enter your Number or \"x\": ")
		scanner.Scan()
		input := scanner.Text()

		// check if user wants to quit enter numbers before 10
		if strings.Compare(input, "x") == 0 {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("That was not an Integer nor x. \n Program ended.\n")
			os.Exit(0)
		}

		intArray = append(intArray, number)
	}
	fmt.Printf("Start Sorting:\n%v\n", intArray)
	bubblesort(intArray)
	fmt.Printf("Array Sorted:\n%v\n", intArray)
}

func bubblesort(intArray []int) {
	for i := len(intArray); i > 1; i-- {
		for pointer := 0; pointer < i-1; pointer++ {
			if intArray[pointer] > intArray[pointer + 1] {
				swap(intArray, pointer)
			}
		}
	}
}

func swap (swapArray []int, pointer int) {
	a := swapArray[pointer]
	swapArray[pointer] = swapArray[pointer + 1]
	swapArray[pointer + 1] = a
}