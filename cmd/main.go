package main

import (
	"flag"
	"fmt"
	"fun-with-go/fun"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	number, err := strconv.Atoi(args[len(args)-2])
	if err != nil {
		fmt.Printf("The given input %v is not a number.\n", err)
		fmt.Println("Exiting the program")
		os.Exit(1)
	}
	fmt.Printf("Printing %d stars in ascending & descending order.\n", number)
	sn := fun.Count{
		N: number,
		B: true,
	}
	sn.Stars()
	sn.B = false
	sn.Stars()

	// wordcount flags
	w := flag.Bool("w", false, "number of words in a file")
	c := flag.Bool("c", false, "number of bytes in a file")
	l := flag.Bool("l", false, "number of lines in a file")
	m := flag.Bool("m", false, "number of characters in a file")

	flag.Parse()

	file := args[len(args)-1]
	fm, err := fun.GetFileMetadata(file)
	if err == nil {
		if *w {
			fmt.Printf("Word count   %d %s\n", fm.WordCount, file)
		}
		if *l {
			fmt.Printf("Line count    %d %s\n", fm.LineCount, file)
		}
		if *c {
			fmt.Printf("Char count  %d %s\n", fm.ByteCount, file)
		}
		if *m {
			fmt.Printf("Rune count  %d %s\n", fm.RuneCount, file)
		}
		if !*c && !*w && !*l && !*m {
			fmt.Printf("    %d   %d  %d %s\n", fm.LineCount, fm.WordCount, fm.ByteCount, file)
		}

	}

}
