package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	stringSlice := readInputFile("./input.txt")
	fmt.Println(stringSlice)
}

// helper function to put the input file into a slice of strings
// (each elements is a line of the txt file)
func readInputFile(path string) []string {
	// var pixelString string
	resultSlice := make([]string, 0)
	absPath, _ := filepath.Abs(path)

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// pixelString = line
		resultSlice = append(resultSlice, line)
	}

	// return pixelString
	return resultSlice
}
