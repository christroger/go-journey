package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food	string
	locomotion	string
	noise	string
}

func (a Animal) Eat() {
	fmt.Printf("This animal eats %s\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("This animals locomotion is %s\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("This animal says %s\n", a.noise)
}

func main() {
	fmt.Printf("Animal Assessment:\nEnter the name of the Animal and what you want to know from it\nE.g. \"cow speak\"\n")
	fmt.Println("Exit the program with typing exit")
	fmt.Println("Animals are: Cow, Snake and Bird")
	fmt.Println("Possible entries are: Speak, Eat, Move")
	var scanner = bufio.NewScanner(os.Stdin)
	
	var animalList = make(map[string]Animal)

	animalList["cow"] = Animal{
		food: "grass",
		locomotion: "walk",
		noise: "moo",
	}

	animalList["bird"] = Animal{
		food: "worms",
		locomotion: "fly",
		noise: "peep",
	}

	animalList["snake"] = Animal{
		food: "mice",
		locomotion: "slither",
		noise: "hsss",
	}

	for true {
		fmt.Printf("> ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		if input == "exit" {
			fmt.Println("Program stopped")
			os.Exit(0)
		}

		query := strings.Fields(input)

		if len(query) > 2 {
			fmt.Println("Invalid Entry, try again")
			continue
		}

		desiredAnimal, empty := animalList[query[0]]
		
		if empty == false {
			fmt.Println("No Matching animal found")
			continue
		}

		switch query[1] {
		case "eat":
			desiredAnimal.Eat()
		case "speak":
			desiredAnimal.Speak()
		case "move":
			desiredAnimal.Move()
		default:
			fmt.Printf("%s is not a valid entry\n", query[1])
		}
	}
}