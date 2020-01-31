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

	minIndex := getMinIndex(layers)

	ones, twos := countOnesAndTwos(layers[minIndex])

	// print solution
	fmt.Println(ones * twos)
}

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

func getMinIndex(layers [][]string) int {
	min, minIndex := 150, 0 // start at 150, max length of one of our nested arrays
	for index, layerSlice := range layers {
		countZeroes := 0
		for _, pixel := range layerSlice {
			if pixel == "0" {
				countZeroes++
			}
		}
		// update min and minIndex if countZeroes is less than the min value
		if countZeroes < min {
			min, minIndex = countZeroes, index
			// fmt.Println(min, minIndex)
		}
	}

	return minIndex
}

func countOnesAndTwos(layer []string) (int, int) {
	ones, twos := 0, 0
	// count ones and twos of the layer with the least zeroes
	for _, pixel := range layer {
		if pixel == "1" {
			ones++
		} else if pixel == "2" {
			twos++
		}
	}
	return ones, twos
}
