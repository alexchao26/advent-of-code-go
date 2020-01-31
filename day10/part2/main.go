package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"./trig"
)

/*
	Overall approach...
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
*/
type Asteroid struct {
	x          int
	y          int
	degOffVert float64
	distance   float64
}

func main() {
	// need to read the input.txt file (one folder up)
	// it will be a 2D slice coming from readInputFile now
	//* actual input file
	stringSlice := readInputFile("../input.txt")

	// !test case
	// stringSlice := readInputFile("./test1.txt")

	// fmt.Println(stringSlice, len(stringSlice))
	gridSlice := make([][]string, len(stringSlice))
	for i, str := range stringSlice {
		gridSlice[i] = strings.Split(str, "")
	}
	// fmt.Println(gridSlice, len(gridSlice))

	//* tests while building the TangetAndDistance function
	// fmt.Println(trig.TangentAndDistance(13, 11, 0, 11))  // 0 13
	// fmt.Println(trig.TangentAndDistance(13, 11, 15, 11)) // 180 2
	// fmt.Println(trig.TangentAndDistance(13, 11, 13, 16)) // 90 5
	// fmt.Println(trig.TangentAndDistance(13, 11, 13, 9))  // 270 2
	// fmt.Println(trig.TangentAndDistance(1, 1, 0, 2)) // 45 some sqrt
	// fmt.Println(trig.TangentAndDistance(1, 1, 2, 2)) // 135 some sqrt
	// fmt.Println(trig.TangentAndDistance(1, 1, 2, 0)) // 225 some sqrt
	// fmt.Println(trig.TangentAndDistance(1, 1, 0, 0)) // 315 some sqrt

	sliceAll := fillSlice(gridSlice)
	// fmt.Println(sliceAll)

	// need to start this just to the left of zero to get that as the first input
	lastDegreeUsed := 359.999999
	// to store the last vaporized asteroid to output its coordinates
	var lastAsteroid Asteroid

	for i := 0; i < 200; i++ {
		// fmt.Println("---lastDegUsed", lastDegreeUsed) // this number should change in very small increments for each loop

		// iterate through all of sliceAll and find the next closest degree
		var indexOfAsteroidToDelete int // will be updated by iMin

		// reset the minDegDiff and minDist for each run of the outer loop
		minDegDiff, minDist := math.Inf(1), math.Inf(1) // I can use inf now b/c I'm using float64's!

		// iterate over the entire slice of asteroids
		for iMin, eAsteroid := range sliceAll {
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
		// remove the element at index indexOfAst..
		// this doesn't maintain order, but I don't care about order right now...
		lastAsteroid = sliceAll[indexOfAsteroidToDelete]
		// swap last element to indexToDelete
		sliceAll[indexOfAsteroidToDelete] = sliceAll[len(sliceAll)-1]
		// re-size slice
		sliceAll = sliceAll[:len(sliceAll)-1]

		// update last deg used by adding the diff to it
		lastDegreeUsed += minDegDiff
		if lastDegreeUsed >= 360 {
			// if we pass over 360, subtract 360
			lastDegreeUsed -= 360
		}
		// fmt.Println(minDegDiff)
		// fmt.Println("lastAsteroid", i, lastAsteroid)
		// fmt.Println(indexOfAsteroidToDelete, sliceAll[indexOfAsteroidToDelete])
	}
	// print the last used asteroid
	fmt.Println(lastAsteroid)

	fmt.Println("advent of code answer: ", lastAsteroid.y*100+lastAsteroid.x)
}

func fillSlice(grid [][]string) []Asteroid {
	result := make([]Asteroid, 0)

	// iterate through the entire grid
	for rowIndex, rowSlice := range grid {
		for colIndex, element := range rowSlice {
			// if an asteroid is found...
			if element == "#" && !(rowIndex == 13 && colIndex == 11) {
				// calculate the degree and dist
				degree, dist := trig.TangentAndDistance(13, 11, rowIndex, colIndex)
				// create an instance of an Asteroid struct and append it to the result slice
				ast := Asteroid{
					x:          rowIndex,
					y:          colIndex,
					degOffVert: degree,
					distance:   dist,
				}
				result = append(result, ast)
			}
		}
	}

	return result
}

func readInputFile(path string) []string {
	// var pixelString string
	pixelSlice := make([]string, 0)
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
		pixelSlice = append(pixelSlice, line)
	}

	// return pixelString
	return pixelSlice
}
