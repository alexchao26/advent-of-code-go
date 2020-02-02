package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// need to read the input.txt file (one folder up)
	// it will be a 2D slice coming from readInputFile now
	stringSlice := readInputFile("../input.txt")
	// fmt.Println(stringSlice, len(stringSlice))
	gridSlice := make([][]string, len(stringSlice))
	for i, str := range stringSlice {
		gridSlice[i] = strings.Split(str, "")
	}
	// fmt.Println(gridSlice, len(gridSlice))

	// will be the final result value
	var result int
	var finalCoords [2]int
	// iterate through entire slice, for each asteroid found, call a helper function
	for rowIndex, rowSlice := range gridSlice {
		for colIndex, element := range rowSlice {
			if element == "#" {
				// # are "asteroids", . are empty space
				// helper function will return how many asteroids are "findable" from the current asteroid

				visibleFromElement := visibleFromAsteroid(gridSlice, rowIndex, colIndex)
				// fmt.Println("visible:", visibleFromElement)

				// take max of return of helper function at end of each loop
				if result < visibleFromElement {
					result = visibleFromElement
					finalCoords[0], finalCoords[1] = rowIndex, colIndex
				}
			}
		}
	}

	// print out the max found
	fmt.Println("result: ", result)
	fmt.Println("best asteroid for the station is at row, col", finalCoords) // [13, 11]
	fmt.Println("from 13, 11 (y, x)", visibleFromAsteroid(gridSlice, 13, 11))
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

// helper function will take, x and y coordinates, and the 2D slices?
// will create a two maps of floats to booleans (one map to cover left side of asteroid, one map to cover right side of asteroid)
//     so that anything that is blocked will not be double counted
// will need edge case handling for planets vertically above or below the current asteroid
func visibleFromAsteroid(grid [][]string, row, col int) (result int) {
	// fmt.Println(grid)
	// make the two maps
	leftMap, rightMap := make(map[float64]bool), make(map[float64]bool)
	// make the two booleans for up and down. nil value is false
	var upBool, downBool bool

	// iterate through every element of the grid slices
	for rowIndex, rowSlice := range grid {
		for colIndex, element := range rowSlice {
			// ensure element is an asteroid & not the asteroid that the helper function is being run on
			if element == "#" && !(row == rowIndex && col == colIndex) {
				rise := rowIndex - row
				run := colIndex - col

				// fmt.Println("avoid zero division")
				// handle if the found asteroid is directly above the inputted row/col asteroid
				if run == 0 {
					if rise < 0 {
						// check down
						// note that up and down are semantically "flipped" due to the 2 row being "above" the 0 row
						if !downBool {
							downBool = true
							result++
						}
					} else {
						if !upBool {
							upBool = true
							result++
						}
					}
				} else {
					slope := float64(rise) / float64(run)
					// fmt.Println(slope)
					// handle left or right map
					if run < 0 {
						// leftMap
						if _, inLeftMap := leftMap[slope]; !inLeftMap {
							leftMap[slope] = true
							result++
						}
					} else {
						// rightMap
						if _, inRightMap := rightMap[slope]; !inRightMap {
							rightMap[slope] = true
							result++
						}
					}
				}
			}
		}
	}

	// fmt.Println(leftMap, rightMap, upBool, downBool)
	return result
}
