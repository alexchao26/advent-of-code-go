package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

func main() {
	// read the input file and put it into a slice for each character
	pixelString := util.ReadFile("../input.txt")
	charsSlice := strings.Split(pixelString, "")

	// make the characers into a 2D slice divided into layers
	layers := makeLayers(charsSlice)

	// for each pixel (all layers combined), iterate through all layers of the pixel
	// iterate from the top layer until a 1 or 0 is found, set it to that value
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

	// Print the six lines (25 characters at a time) individually so the word is legible
	for i := 0; i < 6; i++ {
		fmt.Println(finalString[i*25 : (i+1)*25])
	}
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
