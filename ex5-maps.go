package main

import "fmt"

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and the value
	for key, value := range beatles  {
		fmt.Printf("%s == %s \n", key, value)
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	if bob, ok := beatles["Bob"]; ok {
		fmt.Printf("bob == %s", bob)
	} else {
		fmt.Printf("didnt find bob")
	}
}
