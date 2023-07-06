package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	c := flag.Bool("c", true, "count the bytes of a file")
	flag.Parse()
	fmt.Println("c:", *c)

	f := flag.Arg(0)

	data, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Printf("Number of bytes: %d, %s\n", len(data), f)

	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	countlines := 0
	wordcount := 0
	//charcount := 0
	runecount := 0
	//new := 0
	for {
		line, _, err := reader.ReadLine()
		//fmt.Println(line)
		if err == io.EOF {
			break
		}
		countlines++
		//charcount += utf8.RuneCountInString(string(line))
		//new += utf8.RuneCount(line)
		s := []rune(string(line))
		words := strings.Fields(string(line)) // this gets word count
		wordcount += len(words)
		runecount += len(s) + 1 // adding newline character - is this correct?

	}
	fmt.Printf("Number of lines: %d, %s\n", countlines, f)
	fmt.Printf("Number of words: %d, %s\n", wordcount, f)
	//fmt.Printf("Number of characters: %d, %s\n", charcount, f)
	fmt.Printf("Number of characters: %d, %s\n", runecount, f)
}
