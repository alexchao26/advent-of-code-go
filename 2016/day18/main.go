package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/util"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	var ans int
	if part == 1 {
		ans = likeARogue(util.ReadFile("./input.txt"), 40)
	} else {
		ans = likeARogue(util.ReadFile("./input.txt"), 400000)
	}
	fmt.Println("Output:", ans)
}

// this could be sped up by memoizing rows to their next row, or their counts
// but this is more than fast enough... ~3s to run both parts
func likeARogue(input string, numRows int) int {
	lastRow := "." + input + "."
	patterns := "^^. .^^ ^.. ..^" // to use with strings.Contains

	var actualRows []string
	for len(actualRows) < numRows {
		// add last row
		actualRows = append(actualRows, lastRow[1:len(lastRow)-1])

		// generate the next row
		nextRow := "." // start w/ safe cell in the left wall
		for i := 1; i < len(lastRow)-1; i++ {
			threeAbove := lastRow[i-1 : i+2]
			if strings.Contains(patterns, threeAbove) {
				nextRow += "^"
			} else {
				nextRow += "."
			}
		}
		nextRow += "."
		// assign to last row
		lastRow = nextRow
	}

	// count safe tiles
	var count int
	for _, row := range actualRows {
		for _, v := range row {
			if v == '.' {
				count++
			}
		}
	}

	return count
}
