package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")

	charsSlice := strings.Split(input, "")

	layers := makeLayers(charsSlice)

	minIndex := getMinIndex(layers)

	ones, twos := countOnesAndTwos(layers[minIndex])

	// print solution
	fmt.Println(ones * twos)
}

func makeLayers(charsSlice []string) [][]string {
	layers := make([][]string, 0)
	// layers are 25x6 = 150 characters
	// there are multiple layers of the same size
	for i := 0; i*150 < len(charsSlice); i++ {
		layers = append(layers, charsSlice[150*i:150*(i+1)])
	}

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
		}
	}

	return minIndex
}

func countOnesAndTwos(layer []string) (int, int) {
	ones, twos := 0, 0
	// count ones and twos of the layer with the least zeroes
	for _, pixel := range layer {
		switch pixel {
		case "1":
			ones++
		case "2":
			twos++
		}
	}
	return ones, twos
}
