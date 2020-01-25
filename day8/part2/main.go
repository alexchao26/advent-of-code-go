package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	pixelString := readInputFile("./input.txt")
	// fmt.Println(pixelString)

	charsSlice := strings.Split(pixelString, "")
	// fmt.Println(charsSlice)

	layers := makeLayers(charsSlice)

	// fmt.Println(layers)

	final := make([]string, 0)
	for i := 0; i < 150; i++ {
		final = append(final, "2")
		for _, oneLayer := range layers {
			if oneLayer[i] == "1" {
				final[i] = "0" // zero to view the final image
				break
			} else if oneLayer[i] == "0" {
				final[i] = " " // blank space so it's easier to see the final image
				break
			}
		}
	}

	finalString := strings.Join(final, "")

	line1, line2, line3, line4, line5, line6 := finalString[0:25], finalString[25:50], finalString[50:75], finalString[75:100], finalString[100:125], finalString[125:150]

	// return the final image, flattened
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(line5)
	fmt.Println(line6)
}

// helper functions
func readInputFile(path string) string {
	var pixelString string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pixelString = line
	}

	return pixelString
}

func makeLayers(charsSlice []string) [][]string {
	layers := make([][]string, 0)
	// layers are 25x6 = 150 characters
	// there are multiple layers of the same size
	for i := 0; i*150 < len(charsSlice); i++ {
		layers = append(layers, charsSlice[150*i:150*(i+1)])
	}
	// fmt.Println(layers)
	return layers
}
