package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// Seed our randomizer
	rand.Seed(time.Now().UnixNano())
}

type Command struct {
	ID     int
	Result string
}

// Task: Implement the error interface on the Command type
// When the error is printed, it should output both the ID and Result
func (c Command) Error() string {
	return fmt.Sprintf("%d / %s", c.ID, c.Result)
}

func main() {
	fmt.Println("Starting")
	if err := process(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Completed")
}

// this process will randomly return different errors or no error at all
func process() error {
	i := rand.Int() % 3
	switch i {
	case 0:
		c := Command{ID: 1, Result: "unable to initialize command"}
		return c
	case 1:
		return errors.New("This is a generic error...")
	default:
		return nil
	}
}
