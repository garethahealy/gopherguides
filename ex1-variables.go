package main

import "fmt"

func main() {
	// Declare the following variables with zero values:
	// name: command, type string
	// name: valid, type bool
	var command string = ""
	var valid bool

	// print the value of those variables
	fmt.Println(fmt.Sprintf("command %s", command))
	fmt.Println(fmt.Sprintf("valid %v", valid))

	// declare and initialize the following variables:
	// name: foo, type: int, value: 5
	// name bar, type: bool, value: true
	foo := 5
	bar := true

	// print the value of those variables
	fmt.Println(fmt.Sprintf("foo %d", foo))
	fmt.Println(fmt.Sprintf("bar %v", bar))
}
