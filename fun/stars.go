package fun

import (
	"fmt"
)

// PrintStars print stars in increase/decreasing count
type PrintStars interface {
	Stars()
}

// Count struct to declare n and
// true or false for how to print stars - increasing or decreasing
type Count struct {
	N int
	B bool
}

// Stars function to display ascending or descending count of stars
func (c *Count) Stars() {
	if c.B {
		for i := 1; i <= c.N; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	} else {
		for i := c.N; i >= 1; i-- {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}
