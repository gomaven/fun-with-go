package main

import (
	"fmt"
	"os"
	"strconv"
)

// PrintStars print stars in increase/decreasing count
type PrintStars interface {
	Stars()
}

// Count struct to declare n and
// true or false for how to print stars - increasing or decreasing
type Count struct {
	n int
	b bool
}

// Stars function to display ascending or descending count of stars
func (c *Count) Stars() {
	if c.b {
		for i := 1; i <= c.n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	} else {
		for i := c.n; i >= 1; i-- {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}

func main() {
	//var number int
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("The given input %v is not a number.\n", err)
		fmt.Println("Exiting the program")
		os.Exit(1)
	}
	fmt.Printf("Printing %d stars in ascending & descending order.\n", number)
	c := Count{
		n: number,
		b: true,
	}
	c.Stars()
	c.b = false
	c.Stars()
}
