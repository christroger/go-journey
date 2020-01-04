package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write name of text file that contains names: ")
	fileName, _ := reader.ReadString('\n')
	fileName = fileName[:len(fileName)-1]
	fmt.Println("Input name is: " + fileName)

	dat, _ := ioutil.ReadFile(fileName)
	nameList := string(dat)
	fmt.Println("Read result: " + nameList)
	nameArray := strings.Split(nameList, "\r\n")

	for _, element := range nameArray {
		fullName := strings.Split(element, " ")
		var slice Name
		slice.fname = fullName[0]
		slice.lname = fullName[1]
		fmt.Println(slice.fname + " " + slice.lname)
	}

}
