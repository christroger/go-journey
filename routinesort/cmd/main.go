/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/
package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sync"
	"sort"
)

func main() {
	fmt.Printf("Sorting Assessment:\nEnter at least 8 Numbers spaced by Spaces to Sort:")
	fmt.Printf("Example: \"4 8 5 1 3 2 7 6\"")
	var scanner = bufio.NewScanner(os.Stdin)
	var intArray = make([]int, 0)

	fmt.Println("Please Enter Your Numbers:")
	scanner.Scan()
	input := scanner.Text()

	inputArray := strings.Split(input, " ")

	if len(inputArray) == 0 {
		fmt.Println("You entered an empty string. Therefore no sorting can be done")
		os.Exit(0)
	}
	divide := int((float64(len(inputArray)) / 4.0) + 0.5)

    for i := range inputArray {
		in, _ := strconv.Atoi(inputArray[i])
        intArray = append(intArray, in)
	}

	var intArray1 = intArray[0 : divide]
	var intArray2 = intArray[divide : (divide*2)]
	var intArray3 = intArray[(divide*2) : (divide*3)]
	var intArray4 = intArray[(divide*3) : len(intArray)]
	var wg sync.WaitGroup
	wg.Add(4)
	go sortFunc(intArray1, 1, &wg)
	go sortFunc(intArray2, 2, &wg)
	go sortFunc(intArray3, 3, &wg)
	go sortFunc(intArray4, 4, &wg)
	wg.Wait()

	var sortedArray = make([]int, 0)
	var toMerge = true
	count := 0

	for toMerge {
		mergArray := make([]int, 0)
		if len(intArray1) >= count {
			mergArray = append(mergArray, intArray1[count])
		}
		if len(intArray2) >= count {
			mergArray = append(mergArray, intArray2[count])
		}
		if len(intArray3) >= count {
			mergArray = append(mergArray, intArray3[count])
		}
		if len(intArray4) >= count {
			mergArray = append(mergArray, intArray4[count])
		}
		if len(mergArray) > 0 {
			sort.Ints(mergArray)
			sortedArray = append(sortedArray, mergArray...)
		}

		if len(sortedArray) == len(intArray) {
			toMerge = false
		}
		count++
	}

	fmt.Printf("Sorted Array: %v\n", sortedArray)
}

func sortFunc(sortArray []int, number int64, wg *sync.WaitGroup){
	fmt.Printf("Sorting Algorithm %v will sort %v\n", number, sortArray)
	sort.Ints(sortArray)
	wg.Done()

}