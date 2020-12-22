package main

import (
	"flag"
	"fmt"
	"math"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	ans := signalsAndNoise(util.ReadFile("./input.txt"), part)
	fmt.Println("Output:", ans)
}

func signalsAndNoise(input string, part int) string {
	var grid [][]string
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	var indexMaps []map[string]int
	for col := 0; col < len(grid[0]); col++ {
		indexMaps = append(indexMaps, map[string]int{})
		for row := 0; row < len(grid); row++ {
			char := grid[row][col]
			indexMaps[col][char]++
		}
	}

	var mostVersion string  // part 1
	var leastVersion string // part 2
	for col := 0; col < len(indexMaps); col++ {
		var (
			mostChar  string
			mostLen   int
			leastChar string
			leastLen  int = math.MaxInt32
		)
		for k, count := range indexMaps[col] {
			if count > mostLen {
				mostLen = count
				mostChar = k
			}
			if count < leastLen {
				leastLen = count
				leastChar = k
			}
		}
		mostVersion += mostChar
		leastVersion += leastChar
	}
	if part == 1 {
		return mostVersion
	}
	return leastVersion
}
