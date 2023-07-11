package main

import (
	"fmt"
	"fun-with-go/fun"
	"os"
	"strconv"
)

func main() {
	//var number int
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("The given input %v is not a number.\n", err)
		fmt.Println("Exiting the program")
		os.Exit(1)
	}
	fmt.Printf("Printing %d stars in ascending & descending order.\n", number)
	c := fun.Count{
		N: number,
		B: true,
	}
	c.Stars()
	c.B = false
	c.Stars()

	fm, err := fun.GetFileMetadata("../gutenberg.txt")
	if err == nil {
		fmt.Println(fm)
	}
}
