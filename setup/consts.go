package main

import "fmt"

const (
	foo = "A"
	bar = 'A' // any guesses what this will be?
	bin = 2
)

func main() {
	fmt.Printf("foo is a %T with the value of  %q\n", foo, foo)
	fmt.Printf("bar is a %T with the value of  %v\n", bar, bar)
	fmt.Printf("bin is a %T with the value of  %v\n", bin, bin)

	// And for fun
	fmt.Printf("type: %T value: %v\n", bar, string(bar))

	//  this will fail, as you can't assign a value to a constant
	//foo = "B"
}
