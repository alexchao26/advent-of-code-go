package main

import (
	"github.com/alexchao26/advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("../input.txt")
	lines := strings.Split(input, "\n")

	coords := makeMapOfCoordinates(lines)

	for _, line := range lines {
		uniqueID := hasNoOverlap(coords, line)
		if uniqueID != -1 {
			fmt.Println("Unique Plan ID is:", uniqueID)
		}
	}

}

func makeMapOfCoordinates(lines []string) map[string]int {
	seen := make(map[string]int)
	for _, line := range lines {
		// ID := line[:strings.Index(line, " @")]
		row, _ := strconv.Atoi(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])
		col, _ := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])
		width, _ := strconv.Atoi(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])
		height, _ := strconv.Atoi(line[strings.Index(line, "x")+1:])

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				coords := fmt.Sprintf("%vx%v", row+i, col+j)
				seen[coords]++
			}
		}
	}
	return seen
}

// if cut is unique returns the ID, otherwise -1
func hasNoOverlap(seen map[string]int, line string) int {
	ID, _ := strconv.Atoi(line[1:strings.Index(line, " @")])
	row, _ := strconv.Atoi(line[strings.Index(line, "@")+2 : strings.Index(line, ",")])
	col, _ := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, ":")])
	width, _ := strconv.Atoi(line[strings.Index(line, ":")+2 : strings.Index(line, "x")])
	height, _ := strconv.Atoi(line[strings.Index(line, "x")+1:])

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			coords := fmt.Sprintf("%vx%v", row+i, col+j)
			if seen[coords] != 1 {
				return -1
			}
		}
	}
	return ID
}
