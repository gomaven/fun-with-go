package main

import "fmt"

// func forward(n int) {
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

// func reverse(n int) {
// 	for i := n; i >= 1; i-- {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

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

func (c Count) Stars() {
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
	c := Count{
		n: 10,
		b: true,
	}
	c.Stars()
	c.b = false
	c.Stars()
	// forward(n)
	// reverse(n)
}
