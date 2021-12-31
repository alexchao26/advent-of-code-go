package halp

import (
	"fmt"
	"strings"

	"github.com/alexchao26/advent-of-code-go/mathy"
)

// PrintInfiniteGridStrings supports the type map[[2]int]string, determines the
// bounds of that infinite grid, and consolidates it into a string AND PRINTS IT
// zeroValChar should be one character, and replaces any grid coordinate NOT in
// the infiniteGrid
func PrintInfiniteGridStrings(infiniteGrid map[[2]int]string, zeroValChar string) {
	// get bounds
	var firstRow, lastRow, firstCol, lastCol int
	for coord := range infiniteGrid {
		firstRow = mathy.MinInt(firstRow, coord[0])
		lastRow = mathy.MaxInt(lastRow, coord[0])
		firstCol = mathy.MinInt(firstCol, coord[1])
		lastCol = mathy.MaxInt(lastCol, coord[1])
	}

	var sb strings.Builder
	for r := firstRow; r <= lastRow; r++ {
		for c := firstCol; c <= lastCol; c++ {
			coord := [2]int{r, c}
			if val, ok := infiniteGrid[coord]; ok {
				sb.WriteString(val)
			} else {
				sb.WriteString(zeroValChar)
			}
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

// convertMapBoolsToMapStrings is a helper function for printing map[[2]int]bool
// The return map can be passed into PrintInfiniteGridStrings
func convertMapBoolsToMapStrings(m map[[2]int]bool, trueChar, falseChar string) map[[2]int]string {
	converted := map[[2]int]string{}

	for k, v := range m {
		if v {
			converted[k] = trueChar
		} else {
			converted[k] = falseChar
		}
	}

	return converted
}

// PrintInfiniteGridBools supports the type map[[2]int]bool
func PrintInfiniteGridBools(m map[[2]int]bool, trueChar, falseChar string) {
	mapCoordsToStrings := convertMapBoolsToMapStrings(m, trueChar, falseChar)
	PrintInfiniteGridStrings(mapCoordsToStrings, falseChar)
}
