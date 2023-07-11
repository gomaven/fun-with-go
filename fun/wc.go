package fun

import (
	"bufio"
	"fmt"
	"os"
)

type FileMetadata struct {
	RuneCount int
	ByteCount int
	WordCount int
	LineCount int
}

func counter(s *bufio.Scanner) int {
	count := 0
	for s.Scan() {
		count++
	}
	return count
}

func readFromFile(filename string, sf bufio.SplitFunc) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1, err
	}
	defer file.Close()

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(sf)
	return counter(scanner), nil
}

func checkError(err error) error {
	if err != nil {
		return err
	}
	return nil
}
func GetFileMetadata(filename string) (*FileMetadata, error) {

	rc, err := readFromFile(filename, bufio.ScanRunes)
	if err != nil {
		return nil, err
	}
	lc, err := readFromFile(filename, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	bc, err := readFromFile(filename, bufio.ScanBytes)
	if err != nil {
		return nil, err
	}
	wc, err := readFromFile(filename, bufio.ScanWords)
	if err != nil {
		return nil, err
	}
	return &FileMetadata{
		RuneCount: rc,
		LineCount: lc,
		ByteCount: bc,
		WordCount: wc,
	}, nil
}
