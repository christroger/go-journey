package main

import (
	"encoding/json"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter Name and Address Assessment:\nYou will be asked to enter a Name and a Address\n")
	var name = bufio.NewScanner(os.Stdin)
	var address = bufio.NewScanner(os.Stdin)
	var m = make(map[string]string)
	fmt.Printf("Please enter a Name: ")
	name.Scan()
	m["name"] = name.Text()

	fmt.Printf("Please enter a Address: ")
	address.Scan()
	m["address"] = address.Text()
	mar, _ := json.Marshal(m)
	fmt.Printf("Here is the marshalled Object:\n%s\n", string(mar))
}