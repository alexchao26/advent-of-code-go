package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// read input file, split by the new line, then split each line by commas
	input := util.ReadFile("../input.txt")
	sli := strings.Split(input, "\n")

	split1 := strings.Split(sli[0], ",")
	split2 := strings.Split(sli[1], ",")

	coordsMap1 := makeCoordinatesMap(split1)
	coordsMap2 := makeCoordinatesMap(split2)

	lowestSumOfDistances := 1<<31 - 1

	// iterate over all keys & values in coordsMap1
	for key, value1 := range coordsMap1 {
		// check if the same key is in coordsMap2
		if value2, ok := coordsMap2[key]; ok {
			// update lowestSumOfDistances if applicable
			if lowestSumOfDistances > value1+value2 {
				lowestSumOfDistances = value1 + value2
			}
		}
	}

	fmt.Println("Lowest sum of distances", lowestSumOfDistances)
}

// return a map that has keys of strings and a value of an int of steps to reach the coordinate
func makeCoordinatesMap(directionsSlice []string) map[string]int {
	gridOfCoordinates := map[string]int{}
	prevX := 0
	prevY := 0
	runningLength := 0

	for i := 0; i < len(directionsSlice); i++ {
		// grab the current element
		v := directionsSlice[i]

		// split this element into a slice of runes...
		runeSlice := []rune(v)

		// stores number parsed off of this element
		num, _ := strconv.Atoi(string(runeSlice[1:]))
		// fmt.Println(num)

		// loop from 0 to num and add to the map/gridOfCoordinates
		for num > 0 {
			// on each loop increment the runningLength, decrement num
			runningLength++
			num--
			switch runeSlice[0] {
			case 'R':
				// if going right, increment prevX
				prevX++
			case 'L':
				prevX--
			case 'U':
				prevY++
			case 'D':
				prevY--
			}

			// set `${prevX}x${prevY}` to the map with runningLength as the value
			newCoord := strconv.Itoa(prevX) + "x" + strconv.Itoa(prevY)

			_, ok := gridOfCoordinates[newCoord]
			if ok == false {
				gridOfCoordinates[newCoord] = runningLength
			}
		}
	}
	// fmt.Println(gridOfCoordinates)
	return gridOfCoordinates
}
