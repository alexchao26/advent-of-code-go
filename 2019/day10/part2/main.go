package main

import (
	"fmt"
	"math"
	"strings"

	"adventofcode/2019/day10/part2/trig"
	"adventofcode/util"
)

/*
	Overall approach:
	- need to make a map of some kind
		make it a slice where each element is a struct
			- each struct will contain:
				x
				y
				degOffVert float64 (degrees 0 -> 360)
				distance float64
	- iterate through the slice of structs
		- store the index of the minimum distance

		- remove it from the slice of structs
			- if this is the 200th iteration, store the x and y to return at the end
		- NOTE there are a limited number of asteroids because of the fixed size of the input, so having a O(200 * n) where n is the number of Asteroids, is not a _terrible_ time complexity
*/

// Asteroid data
type Asteroid struct {
	x          int
	y          int
	degOffVert float64
	distance   float64
}

func main() {
	// read input.txt file, split it into a slice of lines
	input := util.ReadFile("../input.txt")
	stringSlice := strings.Split(input, "\n")

	// generate 2D grid of each character from stringSlice
	inputGrid := make([][]string, len(stringSlice))
	for i, str := range stringSlice {
		inputGrid[i] = strings.Split(str, "")
	}

	allAsteroids := makeAsteroidsSlice(inputGrid)

	// need to start this just to the left of zero to get that as the first input
	lastDegreeUsed := 359.999999
	// to store the last vaporized asteroid to output its coordinates
	var lastAsteroid Asteroid

	for i := 0; i < 200; i++ {
		// iterate through all of allAsteroids and find the next closest degree
		var indexOfAsteroidToDelete int // will be updated by iMin

		// reset the minDegDiff and minDist for each run of the outer loop
		minDegDiff, minDist := math.Inf(1), math.Inf(1) // I can use inf now b/c I'm using float64's!

		// iterate over the entire slice of asteroids
		for iMin, eAsteroid := range allAsteroids {
			// calculate the degrees difference
			degDiff := eAsteroid.degOffVert - lastDegreeUsed
			if degDiff <= 0 { // account for the diff passing over zero
				degDiff += 360
			}
			// if this asteroid has a smaller degrees difference, update all min values
			if degDiff < minDegDiff {
				minDist = eAsteroid.distance
				minDegDiff = degDiff
				indexOfAsteroidToDelete = iMin
			} else if degDiff == minDegDiff && minDist > eAsteroid.distance {
				// OR if the degDiff is the same but the distance to 13,11 is smaller
				// update just the minDistance and index for this asteroid
				minDist = eAsteroid.distance
				indexOfAsteroidToDelete = iMin
			}
		}

		// remove the element at index indexOfAst
		// this doesn't maintain order, but I don't care about order right now
		lastAsteroid = allAsteroids[indexOfAsteroidToDelete]

		// swap last element to indexToDelete
		allAsteroids[indexOfAsteroidToDelete] = allAsteroids[len(allAsteroids)-1]
		// re-size slice to effectively pop last element off of slice
		allAsteroids = allAsteroids[:len(allAsteroids)-1]

		// update last deg used by adding the diff to it
		lastDegreeUsed += minDegDiff
		if lastDegreeUsed >= 360 {
			// if we pass over 360, subtract 360
			lastDegreeUsed -= 360
		}
	}

	// print the last used asteroid
	fmt.Println("Last asteroid", lastAsteroid)
	// print the AoC-formatted answer
	fmt.Println("Advent of code answer: ", lastAsteroid.y*100+lastAsteroid.x)
}

func makeAsteroidsSlice(grid [][]string) []Asteroid {
	result := make([]Asteroid, 0)

	// iterate through the entire grid
	for rowIndex, rowSlice := range grid {
		for colIndex, element := range rowSlice {
			// if an asteroid is found...
			if element == "#" && !(rowIndex == 13 && colIndex == 11) {
				// calculate the degree and dist
				// degree, dist := trig.TangentAndDistance(13, 11, rowIndex, colIndex)
				// create an instance of an Asteroid struct and append it to the result slice
				ast := Asteroid{
					x:          rowIndex,
					y:          colIndex,
					degOffVert: trig.AngleOffVertical(13, 11, rowIndex, colIndex),
					distance:   trig.Distance(13, 11, rowIndex, colIndex),
				}
				result = append(result, ast)
			}
		}
	}

	return result
}
