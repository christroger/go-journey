package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {name string}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}

type Bird struct {name string}
func (b Bird) Speak() {
	fmt.Println("peep")
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}

type Snake struct {name string}
func (s Snake) Speak() {
	fmt.Println("hsss")
}
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}

func main() {
	fmt.Printf("Animal Assessment:\nYou can create a animal or Query your existing animals\n")
	fmt.Println("Exit the program with typing exit")
	fmt.Println("You can create following animals and give it a name: Cow, Snake and Bird")
	fmt.Println("Possible Query entries are: Speak, Eat, Move")
	fmt.Println("Create example: newanimal tweety bird")
	fmt.Println("Query example: query tweety eat")
	var scanner = bufio.NewScanner(os.Stdin)
	
	var animalList = make(map[string]Animal)

	for true {
		var a Animal
		fmt.Printf("> ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		if input == "exit" {
			fmt.Println("Program stopped")
			os.Exit(0)
		}

		query := strings.Fields(input)

		if len(query) > 3 {
			fmt.Println("Invalid Entry, try again")
			continue
		}

		switch query[0] {
		case "newanimal":
			name := query[1]
			aType := query[2]
			switch aType {
			case "cow":
				animalList[name] = Cow{name}
			case "bird":
				animalList[name] = Bird{name}
			case "snake":
				animalList[name] = Snake{name}
			}
			fmt.Println("Created it!")
		case "query":
			name := query[1]
			action := query[2]
			desiredAnimal, empty := animalList[name]
		
			if empty == false {
				fmt.Println("No Matching animal found")
				continue
			}

			a = desiredAnimal
			switch action {
			case "speak":
				a.Speak()
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			default:
				fmt.Println("You can only query speak, eat or move")
			}
		default:
			fmt.Println("You need to use the keywords newanimal for creation or query for searching")
		}
	}
}
