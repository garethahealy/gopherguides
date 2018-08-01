package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)
}

func printValue(s *string) {
	// print the pointer value
	fmt.Printf("pointer %s \n", s)
	// print the string value
	fmt.Printf("string %s \n", *s)

	// change the string value
	*s = "Bob"
}
