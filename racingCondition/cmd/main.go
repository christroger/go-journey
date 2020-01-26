package main

import (
	"fmt"
    "time"
)

var count int

func race1() {
	time.Sleep(1 * time.Second)
	count += 6
}

func printrace() {
	time.Sleep(1 * time.Second)
	count -= 2
	fmt.Printf("Printrace: %v\n", count)
}

func main() {
	fmt.Println("This function will create a racing condition.")
	for true {
		count = 0
		fmt.Println("Function Start")
		go race1()
		fmt.Printf("Count: %v\n", count)
		go printrace()
		time.Sleep(2 * time.Second)
	}
}

/**
race1 function will sleep for 1 second and increase the count variable by 6.
printrace function will sleep for 1 second, decrease the count variable by 2 and print it.

My code runs as an infinite loop so it is easier to observe the racing condition.

In my code race1 function is called as a goroutine. Then Count will be printed.

The second function printrace is called as a goroutine aswell and will print count.
After the second call, the program will wait for 2 seconds. Then it starts over again.

This can create a racing condition, because as a developer I don't know when the two routines will be executed. Therefore the outputs of count can change. I could observe on the output of printrace numbers of 4 or -2. These numbers differ because the goroutines get executed at different times. 
**/