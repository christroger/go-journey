package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func genDisplaceFn(acc, inVel, inDis float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5 * acc * (t * t) + inVel*t + inDis
	}
	return fn
}

func main () {
	fmt.Printf("This Programm will calculate displacement from time, aceleration, initial velocity and initial displacement.\n")
	var scanner = bufio.NewScanner(os.Stdin)
	
	fmt.Printf("Enter the Acceleration: ")
	scanner.Scan()
	acceleration, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Your Acceleration cant be convertet to Float64")
		os.Exit(0)
	}

	fmt.Printf("Enter the Initial Velocity: ")
	scanner.Scan()
	inVelocity, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Your Initial Velocity cant be convertet to Float64")
		os.Exit(0)
	}

	fmt.Printf("Enter the Initial Displacement: ")
	scanner.Scan()
	inDisplacement, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Your Initial Displacement cant be convertet to Float64")
		os.Exit(0)
	}

	fmt.Printf("Enter the enter Time: ")
	scanner.Scan()
	time, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Your Time cant be convertet to Float64")
		os.Exit(0)
	}

	GenDis := genDisplaceFn(acceleration, inVelocity, inDisplacement)

	fmt.Printf("Displacement after %v Seconds is: %v\n", time, GenDis(time))
	
}