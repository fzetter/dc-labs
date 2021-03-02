package main

import (
	"fmt"
	"os"
)

var name string

func main() {

	// Variables
	arr := os.Args
	name := ""

	// For
	for _,word := range arr[1:] {
		name = fmt.Sprintf("%v %v", name, word)
	}

	// Print
	fmt.Printf("Hello %s, Welcome to the jungle! \n", name)
}
