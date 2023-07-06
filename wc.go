package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// func WordCount(value string) (int, int) {
// 	// Match non-space character sequences.
// 	re := regexp.MustCompile(`[\S]+`)

// 	// Find all matches and return count.
// 	results := re.FindAllString(value, -1)
// 	//fmt.Printf("%T\n", results)
// 	total := 0
// 	for _, v := range results {
// 		//fmt.Println(i, v)
// 		//total += len(results[i])
// 		total += len(v)
// 	}
// 	return len(results), total
// }

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
		runecount += len(s) + 1 // adding newline character

	}
	fmt.Printf("Number of lines: %d, %s\n", countlines, f)
	fmt.Printf("Number of words: %d, %s\n", wordcount, f)
	//fmt.Printf("Number of characters: %d, %s\n", charcount, f)
	fmt.Printf("Number of characters: %d, %s\n", runecount, f)
	//fmt.Printf("Number of characters: %d, %s\n", new, f)

	// https://amehta.github.io/posts/2019/03/wc-implementation-in-go-lang/
	// scanner := bufio.NewScanner(file)
	// lines, words, characters := 0, 0, 0
	// for scanner.Scan() {
	// 	lines++

	// 	line := scanner.Text()
	// 	characters += len(line)

	// 	splitLines := strings.Split(line, " ")
	// 	words += len(splitLines)
	// }
	// fmt.Printf("%8d%8d%8d %s\n", lines, words, characters, file)
}
