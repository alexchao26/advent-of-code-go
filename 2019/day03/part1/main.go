package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"math"
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

	// the highest safe int32... https://golang.org/pkg/math/
	// start this at the highest possible value so it'll be easily overwriteable
	minDist := math.MaxInt32

	// iterate over all keys & values in coordsMap1
	// if that is also in the other map, check if it has a better manhattan dist
	for key := range coordsMap1 {
		if _, ok := coordsMap2[key]; ok {
			if dist := manhattanDistance(key); dist < minDist {
				// update minDist if applicable
				minDist = dist
			}
		}
	}

	fmt.Println("The lowest Manhattan Distance is", minDist)
}

// make a map where string key represents coordinates
// value is an int for distance of wire to this point
func makeCoordinatesMap(directionsSlice []string) map[string]int {
	gridOfCoordinates := map[string]int{}
	prevX := 0
	prevY := 0
	runningLength := 0

	for i := 0; i < len(directionsSlice); i++ {
		// grab the current element/direction
		v := directionsSlice[i]

		// split this element into a slice of runes...
		runeSlice := []rune(v)

		// stores number parsed off of this element
		num, _ := strconv.Atoi(string(runeSlice[1:len(runeSlice)]))

		// loop from 0 to num and add to the map/gridOfCoordinates
		for num > 0 {
			// on each loop increment the runningLength, decrement num
			runningLength++
			num--
			switch runeSlice[0] {
			case 'R':
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

			if _, ok := gridOfCoordinates[newCoord]; !ok {
				gridOfCoordinates[newCoord] = runningLength
			}
		}
	}

	return gridOfCoordinates
}

func manhattanDistance(coord string) int {
	// parse coordinates off of the passed in string key
	split := strings.Split(coord, "x")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])

	// ensure they're both positive
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}

	return x + y
}
